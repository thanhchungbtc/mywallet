package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/api/middleware"
	"github.com/thanhchungbtc/mywallet/server/app/database"
)

type API struct {
	db *database.DB
}

func New(db *database.DB) *API {
	return &API{db: db}
}

func mustGetLoginId(c *gin.Context) uint {
	auth := c.MustGet(middleware.AUTH_IDENTITY).(map[string]interface{})
	v, _ := strconv.ParseInt(fmt.Sprintf("%v", auth["id"]), 10, 64)

	return uint(v)
}

func (a *API) RegisterRoutes(r gin.IRouter) {
	authMiddleware, err := middleware.AuthMiddleware(a.db)
	if err != nil {
		log.Fatalf("System error %v", err)
	}

	r.POST("/auth/login", authMiddleware.LoginHandler)
	r.POST("/auth/register", a.wrapRegister(authMiddleware))
	r.POST("/auth/refresh-token", authMiddleware.RefreshHandler)

	{
		authGroup := r.Group("")
		authGroup.Use(authMiddleware.MiddlewareFunc())
		authGroup.POST("/auth/verify-token", a.wrapVerifyToken(authMiddleware))

		a.categoryRoutes(authGroup.Group("/categories"))
		a.accountRoutes(authGroup.Group("/accounts"))
		a.expenseRoutes(authGroup.Group("/expenses"))
	}

}
