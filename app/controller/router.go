package controller

import (
	"net/http"
)

type Router interface {
	
}

func NewRouter(todoController TodoController) Router {
	return &router{todoController}
}

func (todoRouter *router) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		todoRouter.todoController.GetTodos(w, r)
	case "POST":
		todoRouter.todoController.InsertTodo(w, r)
	case "PUT":
		todoRouter.todoController.UpdateTodo(w, r)
	case "DELETE":
		todoRouter.todoController.DeleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}