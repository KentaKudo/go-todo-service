package grpc

import (
	"context"

	"github.com/KentaKudo/go-todo-service"
	"github.com/KentaKudo/go-todo-service/internal/pb/service"
)

type Server struct {
	todoService todo.TodoService
}

func NewServer(todoService todo.TodoService) *Server {
	return &Server{todoService}
}

func (s *Server) Create(ctx context.Context, req *service.CreateRequest) (*service.CreateResponse, error) {
	id, err := s.todoService.Create(req.GetTitle(), req.GetDescription())
	if err != nil {
		return nil, handleErr(err)
	}

	return &service.CreateResponse{
		Id: id,
	}, nil
}

var serviceStatusMapping = map[service.Todo_Status]todo.TodoStatus{
	service.TODO_STATUS_UNKNOWN:     todo.TodoStatusCreated,
	service.TODO_STATUS_CREATED:     todo.TodoStatusCreated,
	service.TODO_STATUS_IN_PROGRESS: todo.TodoStatusInProgress,
	service.TODO_STATUS_DONE:        todo.TodoStatusDone,
}

var todoStatusMapping = map[todo.TodoStatus]service.Todo_Status{
	todo.TodoStatusCreated:    service.TODO_STATUS_CREATED,
	todo.TodoStatusInProgress: service.TODO_STATUS_IN_PROGRESS,
	todo.TodoStatusDone:       service.TODO_STATUS_DONE,
}

func (s *Server) Update(ctx context.Context, req *service.UpdateRequest) (*service.UpdateResponse, error) {
	t := todo.Todo{
		ID:          req.GetTodo().GetId(),
		Title:       req.GetTodo().GetTitle(),
		Description: req.GetTodo().GetDescription(),
		Status:      serviceStatusMapping[req.GetTodo().GetStatus()],
	}

	if err := s.todoService.Update(t); err != nil {
		return nil, handleErr(err)
	}

	return &service.UpdateResponse{}, nil
}

func (s *Server) Get(ctx context.Context, req *service.GetRequest) (*service.GetResponse, error) {
	t, err := s.todoService.Get(req.GetId())
	if err != nil {
		return nil, handleErr(err)
	}

	return &service.GetResponse{
		Todo: &service.Todo{
			Id:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      todoStatusMapping[t.Status],
		},
	}, nil
}

func (s *Server) List(ctx context.Context, _ *service.ListRequest) (*service.ListResponse, error) {
	ts, err := s.todoService.List()
	if err != nil {
		return nil, handleErr(err)
	}

	var todos []*service.Todo
	for _, t := range ts {
		todos = append(todos, &service.Todo{
			Id:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      todoStatusMapping[t.Status],
		})
	}

	return &service.ListResponse{
		Todos: todos,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *service.DeleteRequest) (*service.DeleteResponse, error) {
	if err := s.todoService.Delete(req.GetId()); err != nil {
		return nil, handleErr(err)
	}

	return &service.DeleteResponse{}, nil
}

func handleErr(err error) error {
	return err
}
