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
	return &service.CreateResponse{}, nil
}

func (s *Server) Update(ctx context.Context, req *service.UpdateRequest) (*service.UpdateResponse, error) {
	return &service.UpdateResponse{}, nil
}

func (s *Server) Get(ctx context.Context, req *service.GetRequest) (*service.GetResponse, error) {
	return &service.GetResponse{}, nil
}

func (s *Server) List(ctx context.Context, req *service.ListRequest) (*service.ListResponse, error) {
	return &service.ListResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, req *service.DeleteRequest) (*service.DeleteResponse, error) {
	return &service.DeleteResponse{}, nil
}
