package workers

import (
	"fmt"

	"github.com/Gurveer1510/task-scheduler/internal/core"
	"github.com/Gurveer1510/task-scheduler/pkg"
)

type WorkerPool struct {
	NoOfWorkers int
	Ch          <-chan core.Task
}

func NewWorkPool(noOfWorkers int, ch <-chan core.Task) *WorkerPool {
	return &WorkerPool{
		NoOfWorkers: noOfWorkers,
		Ch:          ch,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.NoOfWorkers; i++ {
		go w.Worker(i)
	}
}

func (w *WorkerPool) Worker(id int) {
	fmt.Println("Worker started:", id)

	for task := range w.Ch {
		pkg.SendMsg(task.Payload)
	}

}
