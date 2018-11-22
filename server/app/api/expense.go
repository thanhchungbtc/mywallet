package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
	"github.com/thanhchungbtc/mywallet/server/app/service"
)

func (a *API) expenseRoutes(r gin.IRouter) {
	expService := service.NewExpenseService(a.db)

	r.GET("", a.listExpenses(expService)).
		POST("", a.createExpense(expService)).
		GET("/:id", a.retrieveExpense(expService)).
		DELETE("/:id", a.deleteExpense(expService)).
		PUT("/:id", a.updateExpense(expService))
}

type expenseRequest struct {
	UseDate  time.Time `json:"use_date" binding:"required"`
	Amount   int       `json:"amount" binding:"required"`
	Location string    `json:"location"`
	Memo     string    `json:"memo"`

	AccountID  int `json:"account_id" binding:"required"`
	CategoryID int `json:"category_id" binding:"required"`
}

type expenseResponse struct {
	*model.Expense
}

func newExpenseResponse(expense *model.Expense) *expenseResponse {
	return &expenseResponse{expense}
}

func newExpenseListResponse(objects []*model.Expense) []*expenseResponse {
	res := make([]*expenseResponse, 0)
	for _, c := range objects {
		res = append(res, newExpenseResponse(c))
	}
	return res
}

// retrieveExpense GET /api/expenses/1
func (a *API) retrieveExpense(expService *service.ExpenseService) gin.HandlerFunc {

	return func(c *gin.Context) {
		object, err := expService.FindOne(map[string]interface{}{
			"user_id": mustGetLoginId(c),
			"id":      c.Param("id"),
		})
		if err != nil {
			abortWithError(c, 404, err)
			return
		}
		c.JSON(200, newExpenseResponse(object))
	}
}

// listCategories GET /api/expenses
func (a *API) listExpenses(expService *service.ExpenseService) gin.HandlerFunc {

	return func(c *gin.Context) {
		objects, err := expService.FindAll("user_id = ?", mustGetLoginId(c))
		if err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newExpenseListResponse(objects))
	}
}

// createExpense POST /api/expenses
func (a *API) createExpense(expService *service.ExpenseService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req expenseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abortWithError(c, 400, err)
			return
		}

		object := &model.Expense{
			UseDate:    req.UseDate,
			Amount:     req.Amount,
			AccountID:  req.AccountID,
			CategoryID: req.CategoryID,
			Location:   req.Location,
			Memo:       req.Memo,

			UserID: mustGetLoginId(c),
		}
		if err := expService.Create(object); err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newExpenseResponse(object))
	}
}

// updateExpense PUT /api/expenses/1
func (a *API) updateExpense(expService *service.ExpenseService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req expenseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abortWithError(c, 400, err)
			return
		}

		instance, err := expService.FindOne(map[string]interface{}{
			"id":      c.Param("id"),
			"user_id": mustGetLoginId(c),
		})
		if err != nil {
			abortWithError(c, 400, err)
			return
		}

		instance.UseDate = req.UseDate
		instance.Amount = req.Amount
		instance.AccountID = req.AccountID
		instance.CategoryID = req.CategoryID
		instance.Location = req.Location
		instance.Memo = req.Memo

		if err := expService.Save(instance); err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newExpenseResponse(instance))
	}
}

// deleteExpense DELETE /api/expenses/1
func (a *API) deleteExpense(expService *service.ExpenseService) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := expService.Delete(map[string]interface{}{
			"id":      c.Param("id"),
			"user_id": mustGetLoginId(c),
		}); err != nil {
			abortWithError(c, 400, err)
			return
		}

		c.Status(204)
	}

}
