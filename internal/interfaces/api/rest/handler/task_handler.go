package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gurveer1510/task-scheduler/internal/core"
	"github.com/Gurveer1510/task-scheduler/internal/usecase"
)

type TaskHandler struct {
	TaskService usecase.TaskUseCase
}

func NewTaskHandler(taskService usecase.TaskUseCase) TaskHandler {
	return TaskHandler{
		TaskService: taskService,
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

	log.Print("HANDLER",task)
	newTask, err := h.TaskService.CreateTask(r.Context(), task)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}
