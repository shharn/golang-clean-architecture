// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/shharn/golang-clean-architecture/model"

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *TodoRepository) GetAll() ([]model.Todo, error) {
	ret := _m.Called()

	var r0 []model.Todo
	if rf, ok := ret.Get(0).(func() []model.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Todo)
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