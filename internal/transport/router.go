package transport

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log/slog"
	"my_tastAPI/internal/models"
	"my_tastAPI/internal/services"
	"sync"
)

func NewRouter(data models.DataModel, mu *sync.Mutex, log *slog.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", services.HomePage())

	router.Get("/tasks/{nameTask}", services.GetTaskInfo(data, mu, log))

	router.Post("/tasks/add", services.AddTask(data, mu, log))

	router.Post("/tasks/delete", services.DeleteTask(data, mu, log))

	router.Get("/tasks", services.GetAllTasks(data, mu, log))

	return router

}
