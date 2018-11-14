package serializer

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/internal/app/model"
)

type CategoryRequest struct {
	Name    string `json:"name" binding:"required"`
	Memo    string `json:"memo"`
	Balance int    `json:"balance"`
}

type CategoryResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Memo    string `json:"memo"`
	Balance int    `json:"balance"`
}

func NewCategoryRequest(c *gin.Context) *CategoryRequest {
	return &CategoryRequest{}
}

func (req CategoryRequest) ToCategory() *model.Category {
	return &model.Category{
		Name:    req.Name,
		Memo:    req.Memo,
		Balance: req.Balance,
	}
}

func (req CategoryRequest) UpdateTo(category *model.Category) {
	category.Name = req.Name
	category.Memo = req.Memo
	category.Balance = req.Balance
}

func NewCategoryResponse(category *model.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:      category.ID,
		Name:    category.Name,
		Memo:    category.Memo,
		Balance: category.Balance,
	}
}

func NewCategoryListResponse(categories []*model.Category) []*CategoryResponse {
	res := make([]*CategoryResponse, 0)
	for _, c := range categories {
		res = append(res, NewCategoryResponse(c))
	}
	return res
}
