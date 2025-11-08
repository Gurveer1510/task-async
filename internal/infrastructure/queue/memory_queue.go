package queue

import "github.com/Gurveer1510/task-scheduler/internal/core"

type MemoryQueue struct {
	queue chan core.Task
}

type MemoryQueueInterface interface {
	Enqueue(task core.Task) error
	Dequeue() (core.Task, error)
}

func NewMemoryQueue(bufferSize int) *MemoryQueue {
	return &MemoryQueue{
		queue: make(chan core.Task, bufferSize),
	}
}

func (mq *MemoryQueue) Enqueue(task core.Task) error {
	mq.queue <- task
	return nil
}

func (mq *MemoryQueue) Dequeue() (core.Task, error){
	task := <- mq.queue
	return task, nil
}