package handler

import (
	"encoding/json"
	"testing"

	"github.com/fragmenta/mux/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ErrSystemError = errors.New("Oops...An error occurred")
)

func abortWithError(c *gin.Context, code int, err ...error) {
	log.Printf("%+v", err)
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{"error": err[0].Error()})
		return
	}
	c.AbortWithStatus(code)
}

func abortSystemError(c *gin.Context, err error) {
	abortWithError(c, 500, errors.Wrap(ErrSystemError, err.Error()))
}

func abortUnauthorize(c *gin.Context, err ...error) {
	abortWithError(c, 401, err...)
}

func assertIsJSONFormat(t *testing.T, value interface{}) {
	if _, err := json.Marshal(value); err != nil {
		t.Fatalf("expected json but got %v", value)
	}
}
