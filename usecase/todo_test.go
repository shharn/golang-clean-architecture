package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/shharn/golang-clean-architecture/mocks/repository"
	"github.com/shharn/golang-clean-architecture/model"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodosHappyPath(t *testing.T) {
	mockTodos := []model.Todo{
		0: model.Todo{1, "Content 1", time.Now(), time.Now()},
		1: model.Todo{2, "Content 2", time.Now() ,time.Now()},
	}
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("GetAll").Return(mockTodos, nil)

	result := uc.GetAllTodos()

	assert.Equal(t, mockTodos, result)
	mr.AssertExpectations(t)
}

func TestGetAllTodosWithErrorFromRepository(t *testing.T) {
	emptyTodos := []model.Todo{}
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("GetAll").Return(emptyTodos, errors.New("mock error"))

	result := uc.GetAllTodos()

	assert.Equal(t, emptyTodos, result)
	mr.AssertExpectations(t)
}