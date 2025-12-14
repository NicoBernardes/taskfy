package api

import (
	"github.com/NicoBernardes/taskfy.git/internal/services"
	"github.com/go-chi/chi"
)

type Application struct {
	Router      *chi.Mux
	TaskService services.TaskService
}
