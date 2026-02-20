package main

import (
	"log"
	"net/http"

	"github.com/viniciusdevr/initial-todo-go-for-learning/config"
)

func main() {

	db := config.SetupDB()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
