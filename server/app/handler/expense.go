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
	expenseKey = "mw_expense"
)

type ExpenseHandler struct {
	service *service.Service
}

func NewExpenseHandler(service *service.Service) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

func (h *ExpenseHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("", h.list)
	r.POST("", h.create)
	r.Group("/:id").
		Use(h.withExpenseByID).
		GET("", h.retrieve).
		PUT("", h.update).
		DELETE("", h.delete)
}

// list handles /GET api/categories
func (h *ExpenseHandler) list(c *gin.Context) {
	objects, err := h.service.Expense.All()
	if err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(200, serializer.NewExpenseListResponse(objects))
}

// create handles /POST /api/categories
func (h *ExpenseHandler) create(c *gin.Context) {
	var req serializer.ExpenseRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}

	expense := req.ToExpense()
	if err = h.service.Expense.Save(expense); err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(200, serializer.NewExpenseResponse(expense))
}

// withExpenseByID set expense with 'expenseKey' to the context
func (h *ExpenseHandler) withExpenseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		abortWithError(c, 404, err)
		return
	}
	object, err := h.service.Expense.ByID(id)
	if err != nil {
		abortWithError(c, 404, err)
		return
	}
	c.Set(expenseKey, &object)
}

// update handles /PUT /api/categories/:id
func (h *ExpenseHandler) update(c *gin.Context) {
	expense := h.mustGetExpense(c)

	var req serializer.ExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}

	req.UpdateTo(expense)

	if err := h.service.Expense.Save(expense); err != nil {
		abortWithError(c, 400, err)
		return
	}
	c.JSON(200, serializer.NewExpenseResponse(expense))
}

// delete handles /DELETE /api/categories/:id
func (h *ExpenseHandler) delete(c *gin.Context) {
	expense := h.mustGetExpense(c)
	if err := h.service.Expense.Delete(expense); err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// retrieve handles /GET /api/categories/:id
func (h *ExpenseHandler) retrieve(c *gin.Context) {
	object := h.mustGetExpense(c)
	c.JSON(200, serializer.NewExpenseResponse(object))
}

// mustGetExpense is a helper function
func (h *ExpenseHandler) mustGetExpense(c *gin.Context) *model.Expense {
	object, ok := c.MustGet(expenseKey).(*model.Expense)
	if !ok {
		panic(errors.Wrap(ErrSystemError, fmt.Sprintf("%v is not in the context", expenseKey)))
	}
	return object
}
