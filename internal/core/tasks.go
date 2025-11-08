package core

import (
	"time"
)

type Task struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	JobType   string    `json:"job_type"`
	Payload   Email     `json:"payload"`
	RunAt     time.Time `json:"runAt"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
