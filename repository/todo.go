package repository

import (
	"errors"
	"time"

	"github.com/shharn/golang-clean-architecture/model"
)

type TodoRepository interface {
	Create(model.Todo) error
	Update(model.Todo) error
	Delete(int) error
	GetAll() ([]model.Todo, error)
}

type exampleTodoRepository struct {
	todos map[int]model.Todo
	nextId int
}

func (r *exampleTodoRepository) Create(todo model.Todo) error {
	nextId := r.nextId
	todo.Id = nextId
	r.todos[r.nextId] = todo
	r.nextId = nextId + 1
	return nil
}

func (r *exampleTodoRepository) Update(todo model.Todo) error {
	id := todo.Id
	_, ok := r.todos[id]
	if !ok {
		return errors.New("Has no item to update. Id - " + string(id))
	}
	r.todos[id] = todo
	return nil
}

func (r *exampleTodoRepository) Delete(id int) error {
	_, ok := r.todos[id]
	if !ok {
		return errors.New("Has no item to delete. Id - " + string(id))
	}
	delete(r.todos, id)
	return nil
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
		nextId: 3,
	}
}