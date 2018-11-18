package handler

import (
	"regexp"

	"github.com/jinzhu/gorm"

	"github.com/thanhchungbtc/mywallet/server/app/model"

	"github.com/thanhchungbtc/mywallet/server/app/service"

	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/serializer"
)

var (
	tokenRegexp = regexp.MustCompile(`^(?:T|t)oken (\S+$)`)
	authKey     = "MW_AUTH"
	tokenKey    = "MW_TOKEN"
)

type AuthHandler struct {
	service *service.Service
}

func NewAuthHandler(service *service.Service) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) RegisterRoutes(r gin.IRouter) {
	r.
		POST("/login", h.login).
		POST("/register", h.register)

	sub := r.Group("")
	sub.Use(h.withAuth).
		POST("/verify-token", h.verifyToken).
		POST("/profile", h.profile)

}

func (h *AuthHandler) login(c *gin.Context) {
	var req serializer.LoginRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}

	var tokenStr string
	var user *model.User
	if tokenStr, user, err = h.service.Auth.Login(req.Username, req.Password); err != nil {
		abortWithError(c, 400, err)
		return
	}

	c.JSON(200, serializer.NewAuthResponse(tokenStr, user))
}

func (h *AuthHandler) register(c *gin.Context) {
	var req serializer.RegisterRequest
	var err error
	var tokenStr string
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}

	user := req.ToUser()
	if tokenStr, err = h.service.Auth.Register(user); err != nil {
		abortWithError(c, 400, err)
		return
	}
	c.JSON(200, serializer.NewAuthResponse(tokenStr, user))
}

func (h *AuthHandler) withAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		abortUnauthorize(c)
		return
	}

	matches := tokenRegexp.FindStringSubmatch(authHeader)
	if len(matches) != 2 {
		abortUnauthorize(c)
		return
	}

	tokenStr := matches[1]
	if tokenStr == "" {
		abortUnauthorize(c)
		return
	}

	claims, err := h.service.Auth.ParseToken(tokenStr)

	if err != nil {
		abortUnauthorize(c, err)
		return
	}

	c.Set(authKey, claims)
	c.Set(tokenKey, tokenStr)
}

func (h *AuthHandler) profile(c *gin.Context) {
	_, ok := c.MustGet(authKey).(*service.Claims)
	if !ok {
		abortUnauthorize(c)
		return
	}
	c.JSON(200, gin.H{})
}

func (h *AuthHandler) verifyToken(c *gin.Context) {
	claims, ok := c.MustGet(authKey).(*service.Claims)
	if !ok {
		abortUnauthorize(c)
		return
	}
	tokenStr, _ := c.MustGet(tokenKey).(string)

	c.JSON(200, serializer.NewAuthResponse(tokenStr, &model.User{
		Model:    gorm.Model{ID: claims.User.UserID},
		Username: claims.User.Username,
	}))
}
