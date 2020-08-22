package storage

import (
	"github.com/KentaKudo/go-todo-service"
	"github.com/google/uuid"
)

// XXX: mutex unsafe
type InMemory map[string]todo.Todo

func NewInMemory() InMemory {
	return make(InMemory)
}

func (m InMemory) Create(title, description string) (string, error) {
	id := uuid.New().String()

	newTodo := todo.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      todo.TodoStatusCreated,
	}

	m[id] = newTodo

	return id, nil
}

func (m InMemory) Update(todo todo.Todo) error {
	return nil
}

func (m InMemory) Get(id string) (todo.Todo, error) {
	return todo.Todo{}, nil
}

func (m InMemory) List() ([]todo.Todo, error) {
	return []todo.Todo{}, nil
}

func (m InMemory) Delete(id string) error {
	return nil
}
