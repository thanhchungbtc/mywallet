package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/thanhchungbtc/mywallet/internal/config"

	"github.com/fragmenta/mux/log"

	"github.com/thanhchungbtc/mywallet/internal/database"
	"github.com/thanhchungbtc/mywallet/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUsernameOrPasswordInvalid = errors.New("Username or password is not valid")
)

type Auth interface {
	Login(username, password string) (string, *model.User, error)
	Register(user *model.User) (string, error)
	ParseToken(tokenStr string) (*Claims, error)
}

type auth struct{ *database.DB }

func (s *auth) Login(username, password string) (string, *model.User, error) {
	var user model.User
	if err := s.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("Login failed: %v, %v, %v", username, password, err)
		return "", nil, ErrUsernameOrPasswordInvalid
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Login failed: %v, %v, %v", username, password, err)
		return "", nil, ErrUsernameOrPasswordInvalid
	}
	tokenStr, err := generateJwtToken(&user)
	if tokenStr == "" {
		return "", nil, ErrUsernameOrPasswordInvalid
	}
	return tokenStr, &user, err
}

func (s *auth) Register(user *model.User) (tokenStr string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	// set hashed password
	user.Password = string(hashedPassword)

	tx := s.Begin()
	if err = tx.Save(user).Error; err != nil {
		log.Printf("rollback %v", errors.Wrap(err, "register failed"))
		tx.Rollback()
		return
	}
	tx.Commit()
	tokenStr, err = generateJwtToken(user)
	return
}

func (s *auth) ParseToken(tokenStr string) (*Claims, error) {
	var claims Claims
	p := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}

	_, err := p.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})

	return &claims, err
}

type UserClaims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
}

type Claims struct {
	jwt.StandardClaims
	User UserClaims `json:"user"`
}

func generateJwtToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		User:           UserClaims{Username: user.Username, UserID: user.ID},
		StandardClaims: jwt.StandardClaims{},
	})

	return token.SignedString(config.SecretKey)
}
