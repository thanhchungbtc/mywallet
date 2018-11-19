package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/database"
)

type Router interface {
	RegisterRoutes(r gin.IRouter)
}

type API struct {
	db *database.DB
}

func New(db *database.DB) *API {
	return &API{db: db}
}

func (a *API) RegisterRoutes(r gin.IRouter) {
	s := a.db
	apiRouters := []struct {
		router Router
		path   string
	}{
		{NewAuth2Handler(s), "/auth"},
		{NewAccountHandler(s), "/accounts"},
		{NewCategoryHandler(s), "/categories"},
		{NewExpenseHandler(s), "/expenses"},
	}

	for _, router := range apiRouters {
		router.router.RegisterRoutes(r.Group(router.path))
	}
}
