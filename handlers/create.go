package handlers

import (
	"encoding/json"
	"example/go-todo-api/database"
	"example/go-todo-api/structs"
	"net/http"
)

func Create(w http.ResponseWriter, req *http.Request) {
	var task structs.Task

	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error trying to read body json", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	database.Insert(task)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task created successfully"))
}
