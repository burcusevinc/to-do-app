package service

import (
	"go.mod/model"
	"go.mod/repository"
)

// Service package interface has this method
type ITodoService interface {
	GetAllTasks() ([]*model.Todo, error)
	CreateTask(id int, title string) *model.Todo
}

//Service package struct implements repository interface
type TodoService struct {
	repo repository.ITodoRepository
}

// It returns the service struct
func NewTodoService(repo repository.ITodoRepository) ITodoService {
	return &TodoService{
		repo: repo,
	}
}

// It returns a model array to-do
func (s *TodoService) GetAllTasks() ([]*model.Todo, error) {
	// Using repository's function, it returns all task map
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	// Make a new to-do model slice
	newTasks := make([]*model.Todo, 0, len(tasks))

	// Append this slice new tasks
	for _, task := range tasks {
		newTasks = append(newTasks, task)
	}
	// Return new task array and nil error
	return newTasks, nil
}

// It returns a model to-do
func (s *TodoService) CreateTask(id int, title string) *model.Todo {
	// Create new task struct
	task := &model.Todo{
		Id:    id,
		Title: title,
	}
	// Repository's create task method, uses the task model
	return s.repo.CreateTask(task)
}
