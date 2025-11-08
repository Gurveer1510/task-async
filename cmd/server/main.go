package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gurveer1510/task-scheduler/internal/config"
	"github.com/Gurveer1510/task-scheduler/internal/infrastructure/persistance"
	"github.com/Gurveer1510/task-scheduler/internal/infrastructure/queue"
	"github.com/Gurveer1510/task-scheduler/internal/infrastructure/workers"
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
	queue := queue.NewMemoryQueue(10)
	taskHandler := handler.NewTaskHandler(taskService, *queue)

	router := routes.InitRoutes(&taskHandler)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("PORT not found in config")
	}

	port := cfg.APP_PORT

	workerPool := workers.NewWorkPool(3, taskRepo, *queue)

	workerPool.Start()

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatalf("error in starting the server: %v", err)
	}
}
