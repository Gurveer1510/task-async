package routes

import (
	"github.com/Gurveer1510/task-scheduler/internal/interfaces/api/rest/handler"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(taskHandler *handler.TaskHandler) chi.Router {
	router := chi.NewRouter()
	router.Route("/tasks", func(r chi.Router){
		r.Post("/create", taskHandler.CreateTask)
	})
	return router
}
