package api

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/server/app/model"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Email                string `json:"email" binding:"required,email"`
	Username             string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

func (a *API) wrapRegister(mw *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req registerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			abortWithError(c, 400, err)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			abortWithError(c, 400, err)
			return
		}
		user := &model.User{
			Username: req.Username,
			Password: string(hashedPassword),
			Email:    req.Email,
		}
		if err = a.db.Create(user).Error; err != nil {
			abortWithError(c, 400, err)
			return
		}
		token, expired, err := mw.TokenGenerator(user)
		if err != nil {
			abortWithError(c, 401, err)
			return
		}
		mw.LoginResponse(c, 200, token, expired)
	}
}

func (a *API) wrapVerifyToken(mw *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := a.db.Where("id = ?", mustGetLoginId(c)).First(&user).Error; err != nil {
			abortWithError(c, 400, err)
			return
		}
		token, expired, err := mw.TokenGenerator(user)
		if err != nil {
			abortWithError(c, 401, err)
			return
		}

		mw.LoginResponse(c, 200, token, expired)
	}
}
