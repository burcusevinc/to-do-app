package server

import (
	"fmt"
	"go.mod/handler"
	"go.mod/repository"
	"go.mod/service"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) error {
	repository := repository.NewTodoRepository()
	service := service.NewTodoService(repository)
	handler := handler.NewTodoHandler(service)

	// handle function for given endpoint
	http.HandleFunc("/api/v1/tasks", handler.HandlerMethods)
	// listen given port
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
