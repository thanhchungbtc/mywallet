package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/thanhchungbtc/mywallet/server/app/serializer"
	"github.com/thanhchungbtc/mywallet/server/app/service"

	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

var (
	categoryKey = "mw_category"
)

type CategoryHandler struct {
	service *service.Service
}

func NewCategoryHandler(service *service.Service) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("", h.list)
	r.POST("", h.create)
	r.Group("/:id").
		Use(h.withCategoryByID).
		GET("", h.retrieve).
		PUT("", h.update).
		DELETE("", h.delete)
}

// list handles /GET api/categories
func (h *CategoryHandler) list(c *gin.Context) {
	objects, err := h.service.Category.All()
	if err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(200, serializer.NewCategoryListResponse(objects))
}

// create handles /POST /api/categories
func (h *CategoryHandler) create(c *gin.Context) {
	var req serializer.CategoryRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}

	category := req.ToCategory()
	if err = h.service.Category.Save(category); err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(200, serializer.NewCategoryResponse(category))
}

// withCategoryByID set category with 'categoryKey' to the context
func (h *CategoryHandler) withCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		abortWithError(c, 404, err)
		return
	}
	object, err := h.service.Category.ByID(id)
	if err != nil {
		abortWithError(c, 404, err)
		return
	}
	c.Set(categoryKey, &object)
}

// update handles /PUT /api/categories/:id
func (h *CategoryHandler) update(c *gin.Context) {
	category := h.mustGetCategory(c)

	var req serializer.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}

	req.UpdateTo(category)

	if err := h.service.Category.Save(category); err != nil {
		abortWithError(c, 400, err)
		return
	}
	c.JSON(200, serializer.NewCategoryResponse(category))
}

// delete handles /DELETE /api/categories/:id
func (h *CategoryHandler) delete(c *gin.Context) {
	category := h.mustGetCategory(c)
	if err := h.service.Category.Delete(category); err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// retrieve handles /GET /api/categories/:id
func (h *CategoryHandler) retrieve(c *gin.Context) {
	object := h.mustGetCategory(c)
	c.JSON(200, serializer.NewCategoryResponse(object))
}

// mustGetCategory is a helper function
func (h *CategoryHandler) mustGetCategory(c *gin.Context) *model.Category {
	object, ok := c.MustGet(categoryKey).(*model.Category)
	if !ok {
		panic(errors.Wrap(ErrSystemError, fmt.Sprintf("%v is not in the context", categoryKey)))
	}
	return object
}
