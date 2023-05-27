package handlers

import (
	"encoding/json"
	"example/go-todo-api/database"
	"net/http"
)

func List(w http.ResponseWriter, req *http.Request) {
	tasks := database.List()

	response, _ := json.Marshal(tasks)

	w.Write([]byte(response))
}
