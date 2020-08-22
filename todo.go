package todo

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      TodoStatus
}

type TodoStatus int

const (
	TodoStatusCreated TodoStatus = iota
	TodoStatusInProgress
	TodoStatusDone
)

type TodoService interface {
	Create(string, string) (string, error)
	Update(Todo) error
	Get(string) (Todo, error)
	List() ([]Todo, error)
	Delete(string) error
}
