package repository

import (
	"time"

	"github.com/shharn/golang-clean-architecture/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
}

type exampleTodoRepository struct {
	todos map[int]model.Todo
}

func (r *exampleTodoRepository) GetAll() ([]model.Todo, error) {
	arr := []model.Todo{}
	for _, todo := range r.todos {
		arr = append(arr, todo)
	}
	return arr, nil
}

func NewTodoRepository() TodoRepository {
	return &exampleTodoRepository{
		todos: map[int]model.Todo{
			1: model.Todo{1, "Content 1", time.Now(), time.Now()},
			2: model.Todo{2, "Content 2", time.Now(), time.Now()},
		},
	}
}