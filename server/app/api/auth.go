package api

import (
	"time"

	"github.com/labstack/gommon/log"

	"github.com/thanhchungbtc/mywallet/server/app/serializer"

	"github.com/thanhchungbtc/mywallet/server/app/model"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/config"
)

const (
	identityKey = "auth"
)

type AuthHandler struct {
	service *database.DB
}
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewAuth2Handler(service *database.DB) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) RegisterRoutes(r gin.IRouter) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         config.SecretKey,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: &User{
						ID:       v.ID,
						Username: v.Username,
					},
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims[identityKey]
		},

		Authenticator: func(c *gin.Context) (i interface{}, e error) {
			var req serializer.LoginRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var user *model.User
			if _, user, e = h.service.Auth.Login(req.Username, req.Password); e != nil {
				return "", jwt.ErrMissingLoginValues
			}
			return user, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			return data != nil
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Token",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", h.register)

	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/protect", h.profile)
	}
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

func (h *AuthHandler) profile(c *gin.Context) {
	v, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"user": v,
	})
}
