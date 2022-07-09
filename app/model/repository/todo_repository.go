package repository

import (
	"github.com/yokoiakinori/gosplash-server/app/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(todo entity.TodoEntity) (err error)
}

type todoRepository struct {
	
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func GetTodos() (todos []entity.TodoEntity, err error) {

}

func InsertTodo(todo entity.TodoEntity) (id int, err error) {
	
}

func UpdateTodo(todo entity.TodoEntity) (err error) {
	
}

func DeleteTodo(todo entity.TodoEntity) (err error) {
	
}