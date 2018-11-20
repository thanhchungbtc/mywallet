package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type categoryRequest struct {
	Name string `json:"name" binding:"required"`
	//AvatarURL string `json:"avatar_url" binding:"required"`
	Memo   string `json:"memo"`
	Budget int    `json:"budget"`
}

type categoryResponse struct {
	*model.Category
}

func newCategoryResponse(category *model.Category) *categoryResponse {
	return &categoryResponse{category}
}

func newCategoryListResponse(categories []*model.Category) []*categoryResponse {
	res := make([]*categoryResponse, 0)
	for _, c := range categories {
		res = append(res, newCategoryResponse(c))
	}
	return res
}

// retrieveCategory GET /api/categories/1
func (a *API) retrieveCategory(c *gin.Context) {
	var object model.Category
	userId := mustGetLoginId(c)
	if err := a.db.
		Where(map[string]interface{}{
			"user_id": userId,
			"id":      c.Param("id"),
		}).
		First(&object).Error; err != nil {
		abortWithError(c, 404, err)
		return
	}
	c.JSON(200, newCategoryResponse(&object))
}

// listCategories GET /api/categories
func (a *API) listCategories(c *gin.Context) {
	var objects []*model.Category
	if err := a.db.
		Where("user_id = ?", mustGetLoginId(c)).
		Find(&objects).Error; err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(200, newCategoryListResponse(objects))
}

// createCategory POST /api/categories
func (a *API) createCategory(c *gin.Context) {
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
	if err := a.db.Create(object).Error; err != nil {
		abortWithError(c, 400, err)
		return
	}
	c.JSON(200, newCategoryResponse(object))
}
