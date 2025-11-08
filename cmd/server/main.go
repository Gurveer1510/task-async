package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Gurveer1510/task-scheduler/internal/adaptors/persistance"
	sch "github.com/Gurveer1510/task-scheduler/internal/adaptors/scheduler"
	"github.com/Gurveer1510/task-scheduler/internal/adaptors/workers"
	"github.com/Gurveer1510/task-scheduler/internal/config"
	"github.com/Gurveer1510/task-scheduler/internal/core"
	"github.com/Gurveer1510/task-scheduler/internal/interfaces/api/rest/handler"
	"github.com/Gurveer1510/task-scheduler/internal/interfaces/api/rest/routes"
	"github.com/Gurveer1510/task-scheduler/internal/usecase"
)

func main() {
	db, err := persistance.NewDatabase()
	if err != nil {
		log.Println("error to connect to DB")
		log.Fatal(err)
	}

	taskRepo := persistance.NewTasksRepo(db)
	taskService := usecase.NewTaskUseCase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := routes.InitRoutes(&taskHandler)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("PORT not found in config")
	}

	port := cfg.APP_PORT

	taskChan := make(chan core.Task)

	scheduler := sch.NewScheduler(*db)
	workerPool := workers.NewWorkPool(3, taskChan, taskRepo)

	workerPool.Start()
	go scheduler.RunScheduler(context.Background(), taskChan)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatalf("error in starting the server: %v", err)
	}
}
