package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
	"github.com/thanhchungbtc/mywallet/server/app/service"
)

func (a *API) accountRoutes(r gin.IRouter) {
	catService := service.NewAccountService(a.db)

	r.GET("", a.listAccounts(catService)).
		POST("", a.createAccount(catService)).
		GET("/:id", a.retrieveAccount(catService)).
		DELETE("/:id", a.deleteAccount(catService)).
		PUT("/:id", a.updateAccount(catService))
}

type accountRequest struct {
	Name string `json:"name" binding:"required"`
	//AvatarURL string `json:"avatar_url" binding:"required"`
	Memo    string `json:"memo"`
	Balance int    `json:"budget"`
}

type accountResponse struct {
	*model.Account
}

func newAccountResponse(account *model.Account) *accountResponse {
	return &accountResponse{account}
}

func newAccountListResponse(accounts []*model.Account) []*accountResponse {
	res := make([]*accountResponse, 0)
	for _, c := range accounts {
		res = append(res, newAccountResponse(c))
	}
	return res
}

// retrieveAccount GET /api/accounts/1
func (a *API) retrieveAccount(catService *service.AccountService) gin.HandlerFunc {

	return func(c *gin.Context) {
		object, err := catService.FindOne(map[string]interface{}{
			"user_id": mustGetLoginId(c),
			"id":      c.Param("id"),
		})
		if err != nil {
			abortWithError(c, 404, err)
			return
		}
		c.JSON(200, newAccountResponse(object))
	}
}

// listAccounts GET /api/accounts
func (a *API) listAccounts(catService *service.AccountService) gin.HandlerFunc {

	return func(c *gin.Context) {
		objects, err := catService.FindAll("user_id = ?", mustGetLoginId(c))
		if err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newAccountListResponse(objects))
	}
}

// createAccount POST /api/accounts
func (a *API) createAccount(catService *service.AccountService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req accountRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abortWithError(c, 400, err)
			return
		}

		object := &model.Account{
			Name:    req.Name,
			Balance: req.Balance,
			Memo:    req.Memo,
			UserID:  mustGetLoginId(c),
		}
		if err := catService.Create(object); err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newAccountResponse(object))
	}
}

// updateAccount PUT /api/accounts/1
func (a *API) updateAccount(catService *service.AccountService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req accountRequest
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
		instance.Balance = req.Balance
		instance.Memo = req.Memo

		if err := catService.Save(instance); err != nil {
			abortWithError(c, 400, err)
			return
		}
		c.JSON(200, newAccountResponse(instance))
	}

}

// deleteAccount DELETE /api/accounts/1
func (a *API) deleteAccount(catService *service.AccountService) gin.HandlerFunc {

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
