package handlers

import (
	"encoding/json"
	"example/go-todo-api/database"
	"example/go-todo-api/structs"
	"net/http"
)

func Update(w http.ResponseWriter, req *http.Request) {
	var task structs.Task

	name := req.URL.Query().Get("Name")

	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error trying to read body json", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	id, found := database.Find(name)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	database.Update(id, task)
	w.Write([]byte("Task updated successfully"))
}
