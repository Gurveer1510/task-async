package workers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Gurveer1510/task-scheduler/internal/infrastructure/persistance"
	"github.com/Gurveer1510/task-scheduler/internal/infrastructure/queue"
	"github.com/Gurveer1510/task-scheduler/pkg"
)

type WorkerPool struct {
	NoOfWorkers int
	TaskRepo    persistance.TaskRepo
	queue       queue.MemoryQueue
}

func NewWorkPool(noOfWorkers int, taskRepo persistance.TaskRepo, q queue.MemoryQueue) *WorkerPool {
	return &WorkerPool{
		NoOfWorkers: noOfWorkers,
		TaskRepo:    taskRepo,
		queue:       q,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.NoOfWorkers; i++ {
		go w.Worker(i)
	}
}

func (w *WorkerPool) Worker(id int) {
	fmt.Println("Worker started:", id)

	for {
		task, _ := w.queue.Dequeue()

		now := time.Now().UTC()
		if task.RunAt.After(now) {
			time.Sleep(task.RunAt.Sub(now))
		}

		if err := pkg.SendMsg(task.Payload); err != nil {
			log.Println("Could not send mail ERROR", err)
			w.TaskRepo.MarkAsPending(context.Background(), task)
			continue
		}
		fmt.Println("HERREEE")
		w.TaskRepo.MarkAsDone(context.Background(), task)
	}

}
