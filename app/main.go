package main

import (
	"net/http"

	"github.com/yokoiakinori/gosplash-server/app/controller"
	"github.com/yokoiakinori/gosplash-server/app/model/repository"
)

var todoRepository = repository.NewTodoRepository()

func main() {
	server := http.Server{
		Addr: ":8000",
	}
	http.HandleFunc("/api/", helloworld)
	server.ListenAndServe()
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world.")
}