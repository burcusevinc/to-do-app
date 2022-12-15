package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mod/mock"
	"go.mod/model"
	"go.mod/service"
	"testing"
)

// It checks get all tasks method, it should return to-do array model
func TestTodoService_GetAllTasks(t *testing.T) {
	// Map model
	repoReturn := map[int]*model.Todo{
		1: {
			Id:    1,
			Title: "buy a milk",
		},
	}
	// Mock the repository interface
	repo := mock.NewMockITodoRepository(gomock.NewController(t))

	// Expect that repository get all method return the map model
	repo.EXPECT().
		GetAllTasks().
		Return(repoReturn, nil).
		Times(1)

	// New service method
	service := service.NewTodoService(repo)
	// Get all method
	tasks, err := service.GetAllTasks()

	// Expected task is a array model
	expectedTasks := make([]*model.Todo, 0)

	// Append the array this model
	expectedTasks = append(expectedTasks, &model.Todo{
		Id:    1,
		Title: "buy a milk",
	})

	// Checks expected task model and returned tasks equality
	assert.Equal(t, &expectedTasks, &tasks)
	assert.Nil(t, err)
}

// If task not exist, get all tasks method should return error
func TestTodoService_GetAllTasks_Error(t *testing.T) {
	// Mock the repository interface
	repo := mock.NewMockITodoRepository(gomock.NewController(t))

	// Expect that repository get all method return nil model and an error
	repo.EXPECT().
		GetAllTasks().
		Return(nil, errors.New("Tasks not found! ")).
		Times(1)

	// New service method
	service := service.NewTodoService(repo)
	// Get all method
	tasks, err := service.GetAllTasks()

	// Checks this string and an error
	assert.Equal(t, "Tasks not found! ", err.Error())
	// Assert that task array is nil
	assert.Nil(t, tasks)
}

func TestTodoService_CreateTask(t *testing.T) {
	// to-do model
	repoReturn := model.Todo{
		Id:    1,
		Title: "buy a milk",
	}
	// Mock the repository interface
	repo := mock.NewMockITodoRepository(gomock.NewController(t))

	// Expect that create task method will return to-do model
	repo.EXPECT().
		CreateTask(&repoReturn).
		Return(&repoReturn).
		Times(1)

	// New service method
	service := service.NewTodoService(repo)
	// Create task method
	newTask := service.CreateTask(1, "buy a milk")

	// Checks new task and returned tasks equality
	assert.Equal(t, &repoReturn, newTask)
}
