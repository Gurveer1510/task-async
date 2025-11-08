package persistance

import (
	"context"
	"time"

	"github.com/Gurveer1510/task-scheduler/internal/core"
)

type TaskInterface interface  {
	MarkAsDone(context.Context, core.Task) error
	CreateTask(context.Context, core.Task) (core.Task,error)
}

type TaskRepo struct {
	Db Database
}

func NewTasksRepo(db *Database) TaskRepo {
	return TaskRepo{
		Db: *db,
	}
}

func (tk *TaskRepo) CreateTask(ctx context.Context, task core.Task) (core.Task, error) {
	query := `INSERT INTO tasks(name,job_type,payload,run_at) 
	VALUES($1,$2,$3,$4) RETURNING id, status, created_at;`

	var id int64
	var status string
	var createdAt time.Time
	// fmt.Println(task)
	err := tk.Db.DB.QueryRow(ctx, query, task.Name, task.JobType, task.Payload, task.RunAt).Scan(&id, &status, &createdAt)
	if err != nil {
		return core.Task{}, err
	}
	task.Id = id
	task.Status = status
	task.CreatedAt = createdAt
	return task, nil
}

func (t *TaskRepo) MarkAsDone(ctx context.Context, task core.Task) error {
	query := `UPDATE tasks SET status='done' WHERE id=$1`

	_, err := t.Db.DB.Exec(ctx, query, task.Id)
	return err
}