package repository

import "go.mod/model"

// Repository package interface has this method
type ITodoRepository interface {
	GetAllTasks() (map[int]*model.Todo, error)
	CreateTask(task *model.Todo) *model.Todo
}

// Repository package struct implements model to-do struct map
type TodoRepository struct {
	Tasks map[int]*model.Todo
}

// It returns the repository struct
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		Tasks: map[int]*model.Todo{
			1: {
				Id:    1,
				Title: "buy a water",
			},
		},
	}
}

// It returns the all to-do model array
func (r *TodoRepository) GetAllTasks() (map[int]*model.Todo, error) {
	tasks := r.Tasks
	return tasks, nil
}

// It returns a model to-do
func (r *TodoRepository) CreateTask(task *model.Todo) *model.Todo {
	r.Tasks[task.Id] = &model.Todo{
		Id:    task.Id,
		Title: task.Title,
	}
	return r.Tasks[task.Id]
}
