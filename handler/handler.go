package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/viniciusdevr/initial-todo-go-for-learning/models"
)

type TaskHandler struct {
	DB *sql.DB
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, title, description, done FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows == nil {
		http.Error(w, "No such tasks", http.StatusNoContent)
		return
	}

	routine := []models.Task{}

	for rows.Next() {
		var rout models.Task
		if err := rows.Scan(&rout.Id, &rout.Title, &rout.Description, &rout.Done); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		routine = append(routine, rout)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routine)
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID"+err.Error(), http.StatusBadRequest)
		return
	}

	var rout models.Task

	row := h.DB.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&rout.Id, &rout.Title, &rout.Description, &rout.Done); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rout)
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "JSON Decode error: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.DB.Exec("INSERT INTO tasks (title, description, done) VALUES (?, ?, ?)",
		newTask.Title, newTask.Description, newTask.Done)
	if err != nil {
		http.Error(w, "Insert task error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	idCreated, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Generate ID error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newTask.Id = idCreated

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	result, err := h.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
