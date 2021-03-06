// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/shharn/golang-clean-architecture/model"

// TodoUsecase is an autogenerated mock type for the TodoUsecase type
type TodoUsecase struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: _a0
func (_m *TodoUsecase) CreateTodo(_a0 model.Todo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Todo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodo provides a mock function with given fields: _a0
func (_m *TodoUsecase) DeleteTodo(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTodos provides a mock function with given fields:
func (_m *TodoUsecase) GetAllTodos() []model.Todo {
	ret := _m.Called()

	var r0 []model.Todo
	if rf, ok := ret.Get(0).(func() []model.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Todo)
		}
	}

	return r0
}

// UpdateTodo provides a mock function with given fields: _a0
func (_m *TodoUsecase) UpdateTodo(_a0 model.Todo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Todo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
