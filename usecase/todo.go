package usecase

import (
	"github.com/shharn/golang-clean-architecture/model"
	"github.com/shharn/golang-clean-architecture/repository"
)

type TodoUsecase interface {
	GetAllTodos() []model.Todo
}

type exampleTodoUsecase struct {
	tr repository.TodoRepository
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