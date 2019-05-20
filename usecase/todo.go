package usecase

import (
	"errors"

	"github.com/shharn/golang-clean-architecture/model"
	"github.com/shharn/golang-clean-architecture/repository"
)

var (
	InvalidIdError = errors.New("The Id should not be 0")
)

type TodoUsecase interface {
	CreateTodo(model.Todo) error
	UpdateTodo(model.Todo) error
	DeleteTodo(int) error
	GetAllTodos() []model.Todo
}

type exampleTodoUsecase struct {
	tr repository.TodoRepository
}

func (s *exampleTodoUsecase) CreateTodo(todo model.Todo) error {
	return s.tr.Create(todo)
}

func (s *exampleTodoUsecase) UpdateTodo(todo model.Todo) error {
	if todo.Id < 1 {
		return InvalidIdError
	}
	return s.tr.Update(todo)
}

func (s *exampleTodoUsecase) DeleteTodo(id int) error {
	if id < 1 {
		return InvalidIdError
	}
	return s.tr.Delete(id)
}

func (s *exampleTodoUsecase) GetAllTodos() []model.Todo {
	todos, err := s.tr.GetAll()
	if err != nil {
		return []model.Todo{}
	}
	return todos
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return &exampleTodoUsecase{tr}
}