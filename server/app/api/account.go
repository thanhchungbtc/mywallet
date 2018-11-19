package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type AccountHandler struct {
	db *database.DB
}

func NewAccountHandler(db *database.DB) *AccountHandler {
	return &AccountHandler{db: db}
}

func (h *AccountHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("", h.list)
	r.POST("", h.create)
	r.Group("/:id").
		Use(h.detailCtx).
		GET("", h.retrieve)

}

func (h *AccountHandler) detailCtx(c *gin.Context) {
	var account model.Account
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		abortWithError(c, 400, errors.Wrap(err, "url is not valid"))
		return
	}
	if db := h.db.Where("id = ? ", id).First(&account); db.Error != nil {
		if db.RecordNotFound() {
			abortWithError(c, 404, db.Error)
			return
		}
		abortWithError(c, 500, errors.Wrap(db.Error, "detailCtx: db failed"))
		return
	}
	c.Set("account", &account)
}

// list
func (h *AccountHandler) list(c *gin.Context) {
	var accounts []*model.Account
	if err := h.db.Find(&accounts).Error; err != nil {
		abortWithError(c, 500, err)
		return
	}
	c.JSON(200, newAccountListResponse(accounts))
}

// create
func (h *AccountHandler) create(c *gin.Context) {
	var req accountRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}
	account := req.account()

	if err = h.db.Save(account).Error; err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(201, newAccountResponse(account))
}

func (h *AccountHandler) retrieve(c *gin.Context) {
	account, ok := c.MustGet("account").(*model.Account)
	if !ok {
		abortWithError(c, 500, errors.New("retrieve: cannot retrieve 'account' context"))
		return
	}
	c.JSON(200, newAccountResponse(account))
}

type accountRequest struct {
	Name   string `binding:"required"`
	Memo   string
	Budget int
}

func (req accountRequest) account() *model.Account {
	return &model.Account{
		Name:   req.Name,
		Memo:   req.Memo,
		Budget: req.Budget,
	}
}

type accountResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Memo   string `json:"memo"`
	Budget int    `json:"budget"`
}

func newAccountResponse(account *model.Account) *accountResponse {
	return &accountResponse{
		ID:     account.ID,
		Name:   account.Name,
		Memo:   account.Memo,
		Budget: account.Budget,
	}
}

func newAccountListResponse(accounts []*model.Account) []*accountResponse {
	var res []*accountResponse
	for _, a := range accounts {
		res = append(res, newAccountResponse(a))
	}
	return res
}
