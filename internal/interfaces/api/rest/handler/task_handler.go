package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gurveer1510/task-scheduler/internal/core"
	"github.com/Gurveer1510/task-scheduler/internal/infrastructure/queue"
	"github.com/Gurveer1510/task-scheduler/internal/usecase"
)

type TaskHandler struct {
	TaskService usecase.TaskUseCase
	queue	queue.MemoryQueue
}

func NewTaskHandler(taskService usecase.TaskUseCase, q queue.MemoryQueue) TaskHandler {
	return TaskHandler{
		TaskService: taskService,
		queue: q,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task core.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in decoding the payload %v", err), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	if task.Name == "" || task.JobType == "" {
		http.Error(w, fmt.Sprintf("error in decoding the payload %v", err), http.StatusBadRequest)
		return
	}

	newTask, err := h.TaskService.CreateTask(r.Context(), task)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
		return
	}

	err = h.queue.Enqueue(newTask)	
	if err != nil {
		log.Print("ERROR ENQUEUING:", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}
