package persistance

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Gurveer1510/task-scheduler/internal/core"
)

type Scheduler struct {
	Db Database
}

var ticker = time.NewTicker(2 * time.Second)

var query = `
SELECT 
    id,
    name,
    job_type,
    payload,
    run_at,
    status,
    created_at
FROM tasks
WHERE run_at <= $1 AND status = 'pending'
LIMIT 1;
`

func NewScheduler(db Database) Scheduler {
	return Scheduler{
		Db: db,
	}
}

func (s *Scheduler) RunScheduler(ctx context.Context, ch chan<- core.Task) {
	for {
		select {
		case <-ticker.C:
			var payloadJSON string
			var task core.Task

			row := s.Db.DB.QueryRow(ctx, query, time.Now())

			err := row.Scan(&task.Id, &task.Name, &task.JobType, &payloadJSON, &task.RunAt, &task.Status, &task.CreatedAt)

			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Println("No pending tasks....")
					continue
				} else {
					fmt.Println("error in query", err, "still continuing to maintain flow")
					continue
				}

			}
			if err := json.Unmarshal([]byte(payloadJSON), &task.Payload); err != nil {
				fmt.Println("Bad Payload JSON:", err)
				continue
			}
			fmt.Println("Task found:", task.Name)
			ch <- task
			_, err = s.Db.DB.Exec(ctx, `UPDATE tasks SET status='processing' WHERE id=$1`, task.Id)
			if err != nil {
				fmt.Println("Could not lock the task: TASK ID:", task.Id)
				continue
			}

		case <-ctx.Done():
			fmt.Println("Scheduler stopping")
			return
		}

	}
}
