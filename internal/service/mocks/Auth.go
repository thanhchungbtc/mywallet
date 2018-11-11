// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/thanhchungbtc/mywallet/internal/model"
import service "github.com/thanhchungbtc/mywallet/internal/service"

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// Login provides a mock function with given fields: username, password
func (_m *Auth) Login(username string, password string) (string, *model.User, error) {
	ret := _m.Called(username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *model.User
	if rf, ok := ret.Get(1).(func(string, string) *model.User); ok {
		r1 = rf(username, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.User)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ParseToken provides a mock function with given fields: tokenStr
func (_m *Auth) ParseToken(tokenStr string) (*service.Claims, error) {
	ret := _m.Called(tokenStr)

	var r0 *service.Claims
	if rf, ok := ret.Get(0).(func(string) *service.Claims); ok {
		r0 = rf(tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Claims)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: user
func (_m *Auth) Register(user *model.User) (string, error) {
	ret := _m.Called(user)

	var r0 string
	if rf, ok := ret.Get(0).(func(*model.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
