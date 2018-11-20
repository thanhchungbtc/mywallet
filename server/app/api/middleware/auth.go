package middleware

import (
	"time"

	"github.com/appleboy/gin-jwt"

	"golang.org/x/crypto/bcrypt"

	"github.com/thanhchungbtc/mywallet/server/app/database"

	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

const (
	AUTH_IDENTITY = "BTC_MW_AUTH"
)

var (
	SecretKey = []byte("secret")
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AuthResponseFunc func(c *gin.Context, data interface{})

func AuthMiddleware(db *database.DB) (*jwt.GinJWTMiddleware, error) {

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         SecretKey,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: AUTH_IDENTITY,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					AUTH_IDENTITY: v.ID,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims[AUTH_IDENTITY]
		},

		Authenticator: func(c *gin.Context) (i interface{}, e error) {
			var req LoginRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var user model.User
			if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
				return "", jwt.ErrFailedAuthentication
			}
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			return &user, nil
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
}
