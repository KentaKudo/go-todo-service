package todo

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      Status
}

type Status int

const (
	StatusCreated Status = iota
	StatusInProgress
	StatusDone
)

type Service interface {
	Create(string, string) (string, error)
	Update(Todo) error
	Get(string) (Todo, error)
	List() ([]Todo, error)
	Delete(string) error
}
