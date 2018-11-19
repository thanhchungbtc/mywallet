package api

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/thanhchungbtc/mywallet/server/app/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/database/mocks"
)

// TestAuthHandler_login
func TestAuthHandler_login(t *testing.T) {
	type fields struct {
		service *database.DB
	}
	type args struct {
		c *gin.Context
	}
	type testcase struct {
		fields   fields
		args     args
		w        *httptest.ResponseRecorder
		wantCode int
	}
	tests := map[string]testcase{
		"invalid json": func() testcase {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(`
{
	"username": "foo"
	"password": "bar"
}
`)))
			token, user, err := "token", &model.User{Username: "foo"}, error(nil)
			mockAuth := &mocks.Auth{}
			mockAuth.
				On("Login", "foo", "bar").
				Return(token, user, err)
			mockService := &database.DB{Auth: mockAuth}
			return testcase{
				fields:   fields{service: mockService},
				args:     args{c: c},
				w:        w,
				wantCode: 400,
			}
		}(),

		"success": func() testcase {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(`
{
	"username": "foo",
	"password": "bar"
}
`)))
			token, user, err := "token", &model.User{Username: "foo"}, error(nil)
			mockAuth := &mocks.Auth{}
			mockAuth.
				On("Login", "foo", "bar").
				Return(token, user, err)
			mockService := &database.DB{Auth: mockAuth}
			return testcase{
				fields:   fields{service: mockService},
				args:     args{c: c},
				w:        w,
				wantCode: 200,
			}
		}(),
	}

	for ttname, tt := range tests {
		t.Run(ttname, func(t *testing.T) {
			h := &AuthHandler{
				service: tt.fields.service,
			}
			h.login(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}

// TestAuthHandler_register
func TestAuthHandler_register(t *testing.T) {
	type fields struct {
		service *database.DB
	}
	type args struct {
		c *gin.Context
	}
	type testcase struct {
		fields   fields
		args     args
		w        *httptest.ResponseRecorder
		wantCode int
	}
	tests := map[string]testcase{
		"invalid json": func() testcase {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(`
{
	"username": "foo",
	"email": "foo@bar.com",
	"password": "bar",
	"password_confirmation": "invalid"
}
`)))
			token, err := "random token", error(nil)
			mockAuth := &mocks.Auth{}
			mockAuth.
				On("Register", mock.Anything).
				Return(token, err)
			mockService := &database.DB{Auth: mockAuth}
			return testcase{
				fields:   fields{service: mockService},
				args:     args{c: c},
				w:        w,
				wantCode: 400,
			}
		}(),
		"success": func() testcase {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(`
{
	"username": "foo",
	"email": "foo@bar.com",
	"password": "bar",
	"password_confirmation": "bar"
}
`)))
			token, err := "random token", error(nil)
			mockAuth := &mocks.Auth{}
			mockAuth.
				On("Register", mock.Anything).
				Return(token, err)
			mockService := &database.DB{Auth: mockAuth}
			return testcase{
				fields:   fields{service: mockService},
				args:     args{c: c},
				w:        w,
				wantCode: 200,
			}
		}(),
	}
	for ttname, tt := range tests {
		t.Run(ttname, func(t *testing.T) {
			h := &AuthHandler{
				service: tt.fields.service,
			}
			h.register(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}

func TestAuthHandler_profile(t *testing.T) {
	type fields struct {
		service *database.DB
	}
	type args struct {
		c *gin.Context
	}
	type testcase struct {
		fields   fields
		args     args
		w        *httptest.ResponseRecorder
		wantCode int
	}
	tests := map[string]testcase{
		// TODO: Add test cases.
	}
	for ttname, tt := range tests {
		t.Run(ttname, func(t *testing.T) {
			h := &AuthHandler{
				service: tt.fields.service,
			}
			h.profile(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}
