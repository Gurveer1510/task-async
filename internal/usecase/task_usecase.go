package usecase

import (
	"context"

	"github.com/Gurveer1510/task-scheduler/internal/adaptors/persistance"
	"github.com/Gurveer1510/task-scheduler/internal/core"
)

type TaskUseCase struct{
	TaskRepo persistance.TaskRepo
}

func NewTaskUseCase(taskRepo persistance.TaskRepo) TaskUseCase {
	return TaskUseCase{
		TaskRepo: taskRepo,
	}
}

func (t *TaskUseCase) CreateTask(ctx context.Context, task core.Task) (core.Task, error) {
	return t.TaskRepo.CreateTask(ctx, task)
}

