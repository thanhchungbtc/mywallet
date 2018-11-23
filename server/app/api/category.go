package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
	"github.com/thanhchungbtc/mywallet/server/app/service"
)

func (a *API) categoryRoutes(r gin.IRouter) {
	catService := service.NewCategoryService(a.db)

	r.GET("", a.listCategories(catService)).
		POST("", a.createCategory(catService)).
		GET("/:id", a.retrieveCategory(catService)).
		DELETE("/:id", a.deleteCategory(catService)).
		PUT("/:id", a.updateCategory(catService))
}

type categoryRequest struct {
	Name string `json:"name" binding:"required"`
	//AvatarURL string `json:"avatar_url" binding:"required"`
	Memo   string `json:"memo"`
	Budget int    `json:"budget"`
}

type categoryResponse struct {
	*model.Category
	TotalExpense int `json:"total_expense"`
}

func newCategoryResponse(category *model.Category) *categoryResponse {
	totalExp := 0
	for _, exp := range category.Expenses {
		totalExp += exp.Amount
	}
	res := &categoryResponse{
		Category:     category,
		TotalExpense: totalExp,
	}
	return res
}

func newCategoryListResponse(categories []*model.Category) []*categoryResponse {
	res := make([]*categoryResponse, 0)
	for _, c := range categories {
		res = append(res, newCategoryResponse(c))
	}
	return res
}

// retrieveCategory GET /api/categories/1
func (a *API) retrieveCategory(catService *service.CategoryService) gin.HandlerFunc {

	return func(c *gin.Context) {
		object, err := catService.FindOne(map[string]interface{}{
			"user_id": mustGetLoginId(c),
			"id":      c.Param("id"),
		})
		if err != nil {
			abortWithError(c, 404, err)
			return
		}
		c.JSON(200, newCategoryResponse(object))
	}
}

// listCategories GET /api/categories
func (a *API) listCategories(catService *service.CategoryService) gin.HandlerFunc {

	return func(c *gin.Context) {
		objects, err := catService.FindAll("user_id = ?", mustGetLoginId(c))
		if err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newCategoryListResponse(objects))
	}
}

// createCategory POST /api/categories
func (a *API) createCategory(catService *service.CategoryService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req categoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abortWithError(c, 400, err)
			return
		}

		object := &model.Category{
			Name:   req.Name,
			Budget: req.Budget,
			Memo:   req.Memo,
			UserID: mustGetLoginId(c),
		}
		if err := catService.Create(object); err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newCategoryResponse(object))
	}
}

// updateCategory PUT /api/categories/1
func (a *API) updateCategory(catService *service.CategoryService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req categoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abortWithError(c, 400, err)
			return
		}

		instance, err := catService.FindOne(map[string]interface{}{
			"id":      c.Param("id"),
			"user_id": mustGetLoginId(c),
		})
		if err != nil {
			abortWithError(c, 400, err)
			return
		}

		instance.Name = req.Name
		instance.Budget = req.Budget
		instance.Memo = req.Memo

		if err := catService.Save(instance); err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newCategoryResponse(instance))
	}

}

// deleteCategory DELETE /api/categories/1
func (a *API) deleteCategory(catService *service.CategoryService) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := catService.Delete(map[string]interface{}{
			"id":      c.Param("id"),
			"user_id": mustGetLoginId(c),
		}); err != nil {
			abortWithError(c, 400, err)
			return
		}

		c.Status(204)
	}

}
