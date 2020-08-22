package storage_test

import (
	"testing"

	"github.com/KentaKudo/go-todo-service/internal/storage"
)

func TestDelete(t *testing.T) {
	sut := storage.NewInMemory()

	if err := sut.Delete("unknown id"); err != nil {
		t.Errorf("error returned: %w", err)
	}
}
