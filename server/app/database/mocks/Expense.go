// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/thanhchungbtc/mywallet/server/app/model"

// Expense is an autogenerated mock type for the Expense type
type Expense struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *Expense) All() ([]*model.Expense, error) {
	ret := _m.Called()

	var r0 []*model.Expense
	if rf, ok := ret.Get(0).(func() []*model.Expense); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Expense)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByID provides a mock function with given fields: id
func (_m *Expense) ByID(id int) (*model.Expense, error) {
	ret := _m.Called(id)

	var r0 *model.Expense
	if rf, ok := ret.Get(0).(func(int) *model.Expense); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Expense)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *Expense) Delete(_a0 *model.Expense) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Expense) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: _a0
func (_m *Expense) Save(_a0 *model.Expense) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Expense) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
