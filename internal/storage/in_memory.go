package storage

import (
	"errors"

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

func (m InMemory) Update(new todo.Todo) error {
	old, ok := m[new.ID]
	if !ok {
		return errors.New("not found")
	}

	if new.Title != "" && old.Title != new.Title {
		old.Title = new.Title
	}

	if new.Description != "" && old.Description != new.Description {
		old.Description = new.Description
	}

	if new.Status != todo.TodoStatusCreated && old.Status != new.Status {
		old.Status = new.Status
	}

	m[new.ID] = old

	return nil
}

func (m InMemory) Get(id string) (todo.Todo, error) {
	t, found := m[id]
	if !found {
		return todo.Todo{}, errors.New("not found")
	}

	return t, nil
}

func (m InMemory) List() ([]todo.Todo, error) {
	var allTodos []todo.Todo
	for _, t := range m {
		allTodos = append(allTodos, t)
	}

	return allTodos, nil
}

func (m InMemory) Delete(id string) error {
	delete(m, id)
	return nil
}
