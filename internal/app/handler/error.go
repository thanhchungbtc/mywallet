package handler

import (
	"github.com/fragmenta/mux/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ErrUsernameOrPasswordInvalid = errors.New("Username or password is not valid")
	ErrSystemError               = errors.New("Oops...An error occurred")
)

func abortWithError(c *gin.Context, code int, err error) {
	log.Printf("%+v", err)
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}

func abortSystemError(c *gin.Context, err error) {
	abortWithError(c, 500, errors.Wrap(ErrSystemError, err.Error()))
}
