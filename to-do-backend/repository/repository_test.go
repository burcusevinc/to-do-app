package repository_test

import (
	"github.com/stretchr/testify/assert"
	"go.mod/model"
	"go.mod/repository"
	"testing"
)

// It checks if there is 1 task exist
func TestRepository_GetAllTasks(t *testing.T) {
	// It returns repository struct with 1 model task
	repo := repository.NewTodoRepository()

	// It returns task array
	tasks, err := repo.GetAllTasks()

	// task array length
	length := len(tasks)

	// Assert that task array length equals to 1
	assert.Equal(t, 1, length)
	// Assert that error is nil
	assert.Nil(t, err)
}

// It checks if the method has added the task
func TestTodoRepository_CreateTask(t *testing.T) {
	// It returns repository struct with 1 model task
	repo := repository.NewTodoRepository()

	// Model struct
	Task := model.Todo{
		Id:    2,
		Title: "buy a milk",
	}
	// Get all task
	tasks, err := repo.GetAllTasks()

	//tasks length
	length := len(tasks)
	// Create task method called with task model
	repo.CreateTask(&Task)

	// Checks tasks length equality
	assert.Equal(t, length+1, len(tasks))
	// Assert that error is nil
	assert.Nil(t, err)

}
