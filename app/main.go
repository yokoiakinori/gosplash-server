package main

import (
	"net/http"
	"fmt"
)

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