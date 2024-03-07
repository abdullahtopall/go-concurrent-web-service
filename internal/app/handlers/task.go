package handlers

import (
	"context"
	"encoding/json"
	"golangTestCase/internal/app/models"
	"golangTestCase/internal/pkg/db"
	"golangTestCase/internal/pkg/utils"
	"log"
	"net/http"
)

var workerPool = utils.NewWorker(10)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	workerPool.AddTask(func() {
		err := db.CreateTask(context.Background(), &task)
		if err != nil {
			log.Println("Error creating task:", err)
			http.Error(w, "Error processing task", http.StatusInternalServerError)
			return
		}

		// Görev başarıyla işlendiyse istemciye başarılı bir yanıt gönderilebilir.
		log.Println("Task processed successfully!")
		// Ancak, bu noktada işlemin tamamlanması beklenmeden hemen bir yanıt gönderilir.
		http.Error(w, "Task is being processed", http.StatusAccepted)
	})
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	workerPool.AddTask(func() {
		err := db.UpdateTask(context.Background(), &task)
		if err != nil {
			log.Println("Error updating task:", err)
			http.Error(w, "Error processing task update", http.StatusInternalServerError)
			return
		}

		log.Println("Task updated successfully!")
		http.Error(w, "Task update is being processed", http.StatusAccepted)
	})
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskID int

	workerPool.AddTask(func() {
		err := db.DeleteTask(context.Background(), taskID)
		if err != nil {
			log.Println("Error deleting task:", err)
			http.Error(w, "Error processing task deletion", http.StatusInternalServerError)
			return
		}

		log.Println("Task deleted successfully!")
		http.Error(w, "Task deletion is being processed", http.StatusAccepted)
	})
}
