package main

import (
	"example/go-todo-api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/create-todo", handlers.Create)
	http.HandleFunc("/list-todos", handlers.List)
	http.HandleFunc("/update-todo", handlers.Update)
	http.HandleFunc("/delete-todo", handlers.Delete)

	http.ListenAndServe(":8090", nil)
}
