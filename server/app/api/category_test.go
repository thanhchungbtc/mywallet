package api

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/database/mocks"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// TestCategoryHandler_list
func TestCategoryHandler_list(t *testing.T) {

	type fields struct {
		service *database.DB
	}
	type args struct {
		c *gin.Context
	}
	type testcase struct {
		fields             fields
		args               args
		w                  *httptest.ResponseRecorder
		wantCode           int
		wantContainsString string
	}
	tests := map[string]testcase{
		"error": func() testcase {
			mockCat := &mocks.Category{}
			mockCat.
				On("All").
				Return(nil, errors.New("any error"))

			mockService := &database.DB{Category: mockCat}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			return testcase{
				fields:             fields{service: mockService},
				args:               args{c: c},
				w:                  w,
				wantCode:           400,
				wantContainsString: "error",
			}
		}(),

		"empty list": func() testcase {
			mockCat := &mocks.Category{}
			mockCat.
				On("All").
				Return([]*model.Category{}, nil)

			mockService := &database.DB{
				Category: mockCat,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			return testcase{
				fields:   fields{service: mockService},
				args:     args{c: c},
				w:        w,
				wantCode: 200,
			}
		}(),
		"list of category": func() testcase {
			mockCat := &mocks.Category{}
			mockCat.
				On("All").
				Return([]*model.Category{
					{Model: gorm.Model{ID: 1}, Name: "cat1"},
				}, nil)

			mockService := &database.DB{Category: mockCat}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
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
			h := &CategoryHandler{
				service: tt.fields.service,
			}
			h.list(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
			if tt.wantContainsString != "" {
				assert.Contains(t, tt.w.Body.String(), tt.wantContainsString)
			}
		})
	}
}

// TestCategoryHandler_create
func TestCategoryHandler_create(t *testing.T) {
	type fields struct {
		service *database.DB
	}
	type args struct {
		c *gin.Context
	}
	type testcase struct {
		fields             fields
		args               args
		w                  *httptest.ResponseRecorder
		wantCode           int
		wantContainsString string
	}
	tests := map[string]testcase{
		"bad request": func() testcase {
			mockCat := &mocks.Category{}
			mockCat.
				On("Save", mock.Anything).
				Return(nil)

			mockService := &database.DB{Category: mockCat}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/api/category", bytes.NewBuffer([]byte(`
{
	"name": "foo, 
	"memo": "bar",
	"balance": 100
}`)))
			return testcase{
				fields:             fields{service: mockService},
				args:               args{c: c},
				w:                  w,
				wantCode:           400,
				wantContainsString: "error",
			}
		}(),

		"success": func() testcase {
			mockCat := &mocks.Category{}
			mockCat.
				On("Save", mock.Anything).
				Return(nil)
			mockService := &database.DB{Category: mockCat}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/api/category", bytes.NewBuffer([]byte(`
{
	"name": "foo", 
	"memo": "bar",
	"balance": 100
}`)))
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
			h := &CategoryHandler{
				service: tt.fields.service,
			}
			h.create(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
			assert.Contains(t, tt.w.Body.String(), tt.wantContainsString)
		})
	}
}

func TestCategoryHandler_update(t *testing.T) {
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
			h := &CategoryHandler{
				service: tt.fields.service,
			}
			h.update(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}

func TestCategoryHandler_delete(t *testing.T) {
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
			h := &CategoryHandler{
				service: tt.fields.service,
			}
			h.delete(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}

func TestCategoryHandler_retrieve(t *testing.T) {
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
			h := &CategoryHandler{
				service: tt.fields.service,
			}
			h.retrieve(tt.args.c)
			assertIsJSONFormat(t, tt.w.Body.String())
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}
