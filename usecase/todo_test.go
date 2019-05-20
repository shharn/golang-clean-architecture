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

func TestCreateTodoHappyPath(t *testing.T) {
	mockTodo := model.Todo{Content: "mock content", CreatedAt: time.Now(), ModifiedAt: time.Now()}
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("Create", mockTodo).Return(nil)

	err := uc.CreateTodo(mockTodo)

	assert.Nil(t, err)
	mr.AssertExpectations(t)
}

func TestCreateTodoWithErrorFromRepository(t *testing.T) {
	mockTodo := model.Todo{Content: "mock content", CreatedAt: time.Now(), ModifiedAt: time.Now()}
	mockErr := errors.New("mock error")
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("Create", mockTodo).Return(mockErr)

	err := uc.CreateTodo(mockTodo)

	assert.Equal(t, mockErr, err)
	mr.AssertExpectations(t)
}

func TestUpdateTodoHappyPath(t *testing.T) {
	mockTodo := model.Todo{Id: 1, Content: "mock content", CreatedAt: time.Now(), ModifiedAt: time.Now()}
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("Update", mockTodo).Return(nil)

	err := uc.UpdateTodo(mockTodo)

	assert.Nil(t, err)
	mr.AssertExpectations(t)
}

func TestUpdateTodoWhenTodoHasInvalidId(t *testing.T) {
	mockTodo := model.Todo{Id: 0, Content: "mock content", CreatedAt: time.Now(), ModifiedAt: time.Now()}
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	err := uc.UpdateTodo(mockTodo)

	assert.Equal(t, InvalidIdError, err)
	mr.AssertExpectations(t)
}

func TestUpdateTodoWhenErrorFromRepository(t *testing.T) {
	mockTodo := model.Todo{Id: 11, Content: "mock content", CreatedAt: time.Now(), ModifiedAt: time.Now()}
	mockErr := errors.New("mock error")
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("Update", mockTodo).Return(mockErr)

	err := uc.UpdateTodo(mockTodo)

	assert.Equal(t, mockErr, err)
	mr.AssertExpectations(t)
}

func TestDeleteTodoHappyPath(t *testing.T) {
	mockId := 11
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("Delete", mockId).Return(nil)

	err := uc.DeleteTodo(mockId)

	assert.Nil(t, err)
	mr.AssertExpectations(t)
}

func TestDeleteTodoWhenInvalidIdIsGiven(t *testing.T) {
	mockId := 0
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	err := uc.DeleteTodo(mockId)

	assert.Equal(t, InvalidIdError, err)
	mr.AssertExpectations(t)
}

func TestDeleteTodoWhenErrorFromRepository(t *testing.T) {
	mockId := 33
	mockErr := errors.New("mock error")
	mr := &mocks.TodoRepository{}
	uc := NewTodoUsecase(mr)

	mr.On("Delete", mockId).Return(mockErr)

	err := uc.DeleteTodo(mockId)

	assert.Equal(t, mockErr, err)
	mr.AssertExpectations(t)
}