package serializer

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type ExpenseRequest struct {
	UseDate    time.Time `json:"use_date" binding:"required"`
	Amount     int       `json:"amount" binding:"exists"`
	CategoryID int       `json:"category_id" binding:"required"`
	AccountID  int       `json:"account_id" binding:"required"`
	Location   string    `json:"location"`
	Memo       string    `json:"memo"`
}

type ExpenseResponse struct {
	ID       int       `json:"id"`
	UseDate  time.Time `json:"use_date"`
	Amount   int       `json:"amount"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Account struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Location string `json:"location"`
	Memo     string `json:"memo"`
}

func NewExpenseRequest(c *gin.Context) *ExpenseRequest {
	return &ExpenseRequest{}
}

func (req ExpenseRequest) ToExpense() *model.Expense {
	return &model.Expense{
		UseDate:    req.UseDate,
		Amount:     req.Amount,
		CategoryID: req.CategoryID,
		AccountID:  req.AccountID,
		Location:   req.Location,
		Memo:       req.Memo,
	}
}

func (req ExpenseRequest) UpdateTo(instance *model.Expense) {
	instance.UseDate = req.UseDate
	instance.Amount = req.Amount
	instance.CategoryID = req.CategoryID
	instance.AccountID = req.AccountID
	instance.Location = req.Location
	instance.Memo = req.Memo
}

func NewExpenseResponse(obj *model.Expense) *ExpenseResponse {
	res := &ExpenseResponse{
		ID:       int(obj.ID),
		UseDate:  obj.UseDate,
		Amount:   obj.Amount,
		Location: obj.Location,
		Memo:     obj.Memo,
	}
	res.Category.ID = obj.CategoryID
	res.Category.Name = obj.Category.Name

	res.Account.ID = obj.AccountID
	res.Account.Name = obj.Account.Name

	return res
}

func NewExpenseListResponse(categories []*model.Expense) []*ExpenseResponse {
	res := make([]*ExpenseResponse, 0)
	for _, c := range categories {
		res = append(res, NewExpenseResponse(c))
	}
	return res
}
