package serializer

import (
	"github.com/thanhchungbtc/mywallet/internal/model"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email                string `json:"email" binding:"required,email"`
	Username             string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		UserID   uint   `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

func (r RegisterRequest) ToUser() *model.User {
	return &model.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

func NewAuthResponse(tokenString string, user *model.User) *AuthResponse {
	a := &AuthResponse{
		Token: tokenString,
	}
	a.User.UserID = user.ID
	a.User.Username = user.Username
	return a
}
