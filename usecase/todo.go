package usecase

import (
	"github.com/shharn/golang-clean-architecture/model"
	"github.com/shharn/golang-clean-architecture/repository"
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
	return s.tr.Update(todo)
}

func (s *exampleTodoUsecase) DeleteTodo(id int) error {
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