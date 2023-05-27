package handlers

import (
	"example/go-todo-api/database"
	"net/http"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("Name")

	id, found := database.Find(name)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	database.Delete(id)
	w.Write([]byte("Delete successfully"))
}
