package handler

import (
	"regexp"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thanhchungbtc/mywallet/internal/database"
	"github.com/thanhchungbtc/mywallet/internal/model"
)

var (
	SecretKey        = []byte("secret")
	tokenRegexp      = regexp.MustCompile(`^(?:T|t)oken (\S+$)`)
	getSecretKeyFunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

type AuthHandler struct {
	db *database.DB
}

func NewAuthHandler(db *database.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) RegisterRoutes(r gin.IRouter) {
	r.
		POST("/login", h.login).
		POST("/register", h.register)

}
func (h *AuthHandler) login(c *gin.Context) {
	var req loginRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}
	var user model.User
	if res := h.db.Where("username = ?", req.Username).First(&user); res.Error != nil {
		if res.RecordNotFound() {
			abortWithError(c, 400, ErrUsernameOrPasswordInvalid)
			return
		}
		abortSystemError(c, res.Error)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		abortWithError(c, 401, err)
		return
	}
	response, err := newAuthResponse(&user)
	if err != nil {
		abortWithError(c, 500, err)
		return
	}

	c.JSON(200, response)
}

func (h *AuthHandler) register(c *gin.Context) {
	var req registerRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		abortWithError(c, 400, err)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		abortWithError(c, 500, errors.Wrap(ErrSystemError, err.Error()))
		return
	}
	user := req.user(string(hashedPassword))

	if err = h.db.Save(user).Error; err != nil {
		abortWithError(c, 500, errors.Wrap(ErrSystemError, err.Error()))
		return
	}

	response, err := newAuthResponse(user)
	if err != nil {
		abortWithError(c, 500, errors.Wrap(ErrSystemError, err.Error()))
		return
	}

	c.JSON(200, response)
}

func generateJwtToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		User:           userResponse{Username: user.Username, UserID: user.ID},
		StandardClaims: jwt.StandardClaims{},
	})

	return token.SignedString(SecretKey)
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerRequest struct {
	Email                string `json:"email" binding:"required,email"`
	Username             string `json:"username" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

func (r registerRequest) user(hashedPassword string) *model.User {
	return &model.User{
		Username: r.Username,
		Email:    r.Email,
		Password: hashedPassword,
	}
}

type userResponse struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
}
type claims struct {
	jwt.StandardClaims
	User userResponse `json:"user"`
}

type authResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

func newAuthResponse(user *model.User) (*authResponse, error) {
	tokenStr, err := generateJwtToken(user)
	if err != nil {
		return nil, err
	}
	return &authResponse{
		Token: tokenStr,
		User:  userResponse{Username: user.Username, UserID: user.ID},
	}, nil
}
