package main

import (
	"net/http"
	"os"

	"github.com/yokoiakinori/gosplash-server/app/controller"
	"github.com/yokoiakinori/gosplash-server/app/model/repository"
	"github.com/joho/godotenv"
)

var todoRepository = repository.NewTodoRepository()
var todoController = controller.NewTodoController(todoRepository)
var todoRouter = controller.NewRouter(todoController)

func main() {
	server := http.Server{
		Addr: ":8000",
	}
	http.HandleFunc("/api/", loadEnv)
	server.ListenAndServe()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込みに失敗しました。")
	}

	message := os.Getenv("TEST")
	fmt.Printf(message)
}