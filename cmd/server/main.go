package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Gurveer1510/task-scheduler/internal/adaptors/persistance"
	"github.com/Gurveer1510/task-scheduler/internal/config"
	"github.com/Gurveer1510/task-scheduler/internal/core"
	"github.com/Gurveer1510/task-scheduler/internal/interfaces/api/rest/handler"
	"github.com/Gurveer1510/task-scheduler/internal/interfaces/api/rest/routes"
	"github.com/Gurveer1510/task-scheduler/internal/usecase"
	"github.com/Gurveer1510/task-scheduler/pkg"
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

	scheduler := persistance.NewScheduler(*db)

for i := range 3 {
	go worker(i, taskChan)
}

go scheduler.RunScheduler(context.Background(), taskChan)
err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
if err != nil {
	log.Fatalf("error in starting the server: %v", err)
}

}

// worker processes tasks from the channel
func worker(id int, ch <-chan core.Task) {
	for task := range ch {
		pkg.SendMsg(task.Payload)
	}
}
