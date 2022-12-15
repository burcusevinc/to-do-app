package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mod/handler"
	"go.mod/mock"
	"go.mod/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

// It checks service returned tasks and handler response equality
func TestTodoHandler_GetAllTasks(t *testing.T) {
	// Mock the service interface
	service := mock.NewMockITodoService(gomock.NewController(t))
	// Returned to-do model
	serviceReturn := []*model.Todo{
		{
			Id:    0,
			Title: "",
		},
	}
	// Expect that service get all method return serviceReturn model and nil error
	service.EXPECT().
		GetAllTasks().
		Return(serviceReturn, nil).
		Times(1)

	// New handler method
	handler := handler.NewTodoHandler(service)
	// Test get request
	req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", http.NoBody)
	// Test response
	res := httptest.NewRecorder()
	// Handler get all method using test request and response
	handler.GetAllTasks(res, req)

	// Expected task is a array model
	expectedTasks := []*model.Todo{{
		Id:    0,
		Title: "",
	}}

	// Json parsed the response's body and give it to struct
	json.Unmarshal(res.Body.Bytes(), expectedTasks)

	// Checks expected task model and returned tasks equality
	assert.Equal(t, serviceReturn, expectedTasks)
	// Checks response result status code equals to 200
	assert.Equal(t, res.Result().StatusCode, http.StatusOK)
	// Checks response content type is "application/json"
	assert.Equal(t, "application/json; charset=UTF-8", res.Header().Get("content-type"))

}

// If service returns error, handler response should return error
func TestTodoHandler_GetAllTasks_Error(t *testing.T) {
	// Mock the service interface
	service := mock.NewMockITodoService(gomock.NewController(t))

	// New error
	serviceError := errors.New("test error")

	// Expect that service get all task method return nil and an error
	service.EXPECT().
		GetAllTasks().
		Return(nil, serviceError).
		Times(1)

	// New handler method
	handler := handler.NewTodoHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", http.NoBody)
	res := httptest.NewRecorder()
	handler.GetAllTasks(res, req)

	// Checks response result status code equals to 500
	assert.Equal(t, res.Result().StatusCode, http.StatusInternalServerError)
	// Checks response body equals to an error
	assert.Equal(t, res.Body.String(), serviceError.Error())

}

// If create method works correctly returns http status created method
func TestTodoHandler_CreateTask(t *testing.T) {
	// Mock the service interface
	service := mock.NewMockITodoService(gomock.NewController(t))

	// Returned to-do model
	serviceReturn := model.Todo{
		Id:    1,
		Title: "buy a milk",
	}
	// Expect that service create method return serviceReturn model
	service.EXPECT().
		CreateTask(1, "buy a milk").
		Return(&serviceReturn).
		Times(1)

	// New handler method
	handler := handler.NewTodoHandler(service)

	// Json encoding of serviceReturn
	body, _ := json.Marshal(serviceReturn)

	// Test request Post method
	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewReader(body))
	// Test response
	res := httptest.NewRecorder()

	// Hanler create task method
	handler.CreateTask(res, req)

	// Checks response result status code equals to 201
	assert.Equal(t, http.StatusCreated, res.Result().StatusCode)
}

func TestTodoHandler_CreateTask_ReturnError(t *testing.T) {
	// Mock service interface
	service := mock.NewMockITodoService(gomock.NewController(t))

	// New handler method
	handler := handler.NewTodoHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", errReader(0))
	res := httptest.NewRecorder()

	handler.CreateTask(res, req)

	// Checks response result status code equals to 500
	assert.Equal(t, res.Result().StatusCode, http.StatusInternalServerError)
}

// Handler Request Methods Test
func TestTodoHandler_HandlerMethods(t *testing.T) {
	t.Run("get all tasks get method", func(t *testing.T) {
		// Mocked service interface
		mockService := mock.NewMockITodoService(gomock.NewController(t))

		// Model to-do array
		serviceReturn := []*model.Todo{
			{
				Id:    0,
				Title: "",
			},
		}

		// Expect that service get all method returns to-do array
		mockService.EXPECT().
			GetAllTasks().
			Return(serviceReturn, nil).
			Times(1)

		handler := handler.NewTodoHandler(mockService)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
		res := httptest.NewRecorder()
		handler.HandlerMethods(res, req)

		// Empty array expected
		expectedTasks := []*model.Todo{{}}
		// Json parsed the response's body and give it to model
		json.Unmarshal(res.Body.Bytes(), expectedTasks)

		// Checks expected task model and returned tasks equality
		assert.Equal(t, serviceReturn, expectedTasks)
		// Checks response result status code equals to 200
		assert.Equal(t, res.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json; charset=UTF-8", res.Header().Get("content-type"))
	})

	t.Run("create task post method", func(t *testing.T) {
		// Mocked service interface
		service := mock.NewMockITodoService(gomock.NewController(t))
		// Model to-do created
		serviceReturn := model.Todo{
			Id:    1,
			Title: "buy a milk",
		}
		// Expect that service create method returns to-do model
		service.EXPECT().
			CreateTask(1, "buy a milk").
			Return(&serviceReturn).
			Times(1)

		handler := handler.NewTodoHandler(service)

		// Encoding Json of serviceReturn model
		body, _ := json.Marshal(serviceReturn)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		res := httptest.NewRecorder()

		handler.HandlerMethods(res, req)

		// Checks response result status code equals to 201
		assert.Equal(t, http.StatusCreated, res.Result().StatusCode)
	})

	t.Run("default method", func(t *testing.T) {
		// Mock service interface
		service := mock.NewMockITodoService(gomock.NewController(t))

		handler := handler.NewTodoHandler(service)
		// Test request put method
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		// Test response
		res := httptest.NewRecorder()
		handler.HandlerMethods(res, req)

		// empty to-do model
		expectedTasks := model.Todo{}
		// Json parsed the response's body and give it to model
		json.Unmarshal(res.Body.Bytes(), expectedTasks)

		// Checks response result status code equals to 405
		assert.Equal(t, res.Result().StatusCode, http.StatusMethodNotAllowed)
		// Checks response body equals to "method not allowed"
		assert.Equal(t, res.Body.Bytes(), []byte("method not allowed"))
	})
}

type errReader int

// Created new error
func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
