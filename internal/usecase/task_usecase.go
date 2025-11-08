package usecase

import (
	"context"
	"log"

	"github.com/Gurveer1510/task-scheduler/internal/adaptors/persistance"
	"github.com/Gurveer1510/task-scheduler/internal/core"
)

type TaskUseCase struct {
	TaskRepo persistance.TaskRepo
}

func NewTaskUseCase(taskRepo persistance.TaskRepo) TaskUseCase {
	return TaskUseCase{
		TaskRepo: taskRepo,
	}
}

func (t *TaskUseCase) CreateTask(ctx context.Context, task core.Task) (core.Task, error) {
	log.Println("USECASE:", task)
	task.RunAt = task.RunAt.UTC()
	log.Println("USECASE:", task)
	return t.TaskRepo.CreateTask(ctx, task)
}
