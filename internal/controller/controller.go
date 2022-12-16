package controller

import (
	"backend-hagowagonetka/internal/services"

	"github.com/go-chi/chi/v5"
)

type HTTPController struct {
	Services *services.Services
}

func NewHTTPController(
	router *chi.Mux,
	Services *services.Services,
) *HTTPController {

	api := chi.NewRouter()
	router.Mount("/api/v1", api)

	ctrl := &HTTPController{
		Services: Services,
	}

	ctrl.register(api)
	return ctrl
}
