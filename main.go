package main

import (
	"log"
	"net/http"

	"github.com/viniciusdevr/initial-todo-go-for-learning/config"
	"github.com/viniciusdevr/initial-todo-go-for-learning/handler"
)

func main() {

	db := config.SetupDB()
	defer db.Close()

	handler := &handler.TaskHandler{
		DB: db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /task", handler.AddTask)
	mux.HandleFunc("GET /task/{id}", handler.GetTaskById)
	mux.HandleFunc("GET /task", handler.GetAllTasks)
	mux.HandleFunc("DELETE /task/{id}", handler.DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
