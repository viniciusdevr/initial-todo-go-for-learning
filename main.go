package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/viniciusdevr/initial-todo-go-for-learning/config"
	"github.com/viniciusdevr/initial-todo-go-for-learning/models"
)

func tasksNotDone(done bool, db *sql.DB) ([]models.Task, error) {

	var routine []models.Task

	rows, err := db.Query("SELECT * FROM tasks WHERE done = ?", done)
	if err != nil {
		return nil, fmt.Errorf("tasksNotDone: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var rout models.Task
		if err := rows.Scan(&rout.Id, &rout.Title, &rout.Description, &rout.Done); err != nil {
			return nil, fmt.Errorf("tasksNotDone: %v", err)
		}
		routine = append(routine, rout)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("tasksNotDone: ", err)
	}

	return routine, nil
}

func taskById(id int, db *sql.DB) (models.Task, error) {
	var rout models.Task

	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&rout.Id, &rout.Title, &rout.Description, &rout.Done); err != nil {
		if err == sql.ErrNoRows {
			return rout, fmt.Errorf("taskById %d: no such task", id)
		}
		return rout, fmt.Errorf("taskById %d: %v", id, err)
	}
	return rout, nil
}

func main() {

	db := config.SetupDB()
	defer db.Close()

	// routine, err := tasksNotDone(false, db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Routine found: %v\n", routine)

	rout, err := taskById(2, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task found: %v\n", rout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
