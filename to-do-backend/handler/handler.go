package handler

import (
	"encoding/json"
	"go.mod/model"
	"go.mod/service"
	"io/ioutil"
	"net/http"
)

// Handler package interface has this method
type ITodoHandler interface {
	GetAllTasks(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	HandlerMethods(w http.ResponseWriter, r *http.Request)
}

// Handler package struct implements service interface
type TodoHandler struct {
	service service.ITodoService
}

// It returns the handler struct
func NewTodoHandler(s service.ITodoService) ITodoHandler {
	return &TodoHandler{service: s}
}

func (h *TodoHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	// Service get all method it returns all tasks
	tasks, err := h.service.GetAllTasks()
	// If it returns nil, return 500 status code and error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //500 HTTP response
		w.Write([]byte(err.Error()))
		return
	}

	// Json encoding of task array
	jsonBytes, _ := json.Marshal(tasks)

	// Added HTTP Header values
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	// 200 Status Code
	w.WriteHeader(http.StatusOK)
	// HTTP reply, encoding tasks
	w.Write(jsonBytes)
}

func (h *TodoHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Empty model to-do
	newTask := model.Todo{}

	// Read the request body
	read, err := ioutil.ReadAll(r.Body)
	// If request body returns error return 500 status code and an error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	// Parses the read json data and stores it in newtask model
	json.Unmarshal(read, &newTask)

	// Service's create task method with newtask model
	h.service.CreateTask(newTask.Id, newTask.Title)

	// Returns 201 status code
	w.WriteHeader(http.StatusCreated)

}

func (h *TodoHandler) HandlerMethods(w http.ResponseWriter, r *http.Request) {
	// method switch-case
	switch {
	// if request method is get, return get all task method
	case r.Method == "GET":
		h.GetAllTasks(w, r)
		return
	case r.Method == "POST":
		h.CreateTask(w, r)
		return
	default:
		// default 405 HTTP status code
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
