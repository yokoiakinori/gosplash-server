package controller

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/yokoiakinori/gosplash-server/app/controller/dto"
	"github.com/yokoiakinori/gosplash-server/app/model/entity"
	"github.com/yokoiakinori/gosplash-server/app/model/repository"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	InsertTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, *http.Request) {
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}
}