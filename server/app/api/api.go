package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/appleboy/gin-jwt"
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
	claims := jwt.ExtractClaims(c)
	v, _ := strconv.ParseInt(fmt.Sprintf("%v", claims[middleware.AUTH_IDENTITY]), 10, 64)

	return uint(v)
}

func (a *API) RegisterRoutes(r gin.IRouter) {
	authMiddleware, err := middleware.AuthMiddleware(a.db)
	if err != nil {
		log.Fatalf("System error %v", err)
	}

	r.POST("/auth/login", authMiddleware.LoginHandler)
	r.POST("/auth/register", a.wrapRegister(authMiddleware))

	r.
		Use(authMiddleware.MiddlewareFunc()).
		POST("/auth/verify-token", a.verifyToken).
		GET("/categories", a.listCategories).
		POST("/categories", a.createCategory).
		GET("/categories/:id", a.retrieveCategory)

}
