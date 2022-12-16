package app

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/controller"
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

func Launch() {
	cfg := config.Get()

	router := chi.NewMux()
	router.Use(middleware.Logger)

	repo := repository.NewRepository(repository.Source{
		User:         cfg.Env.DatabaseUser,
		Password:     cfg.Env.DatabasePassword,
		Host:         cfg.Env.DatabaseHost,
		Port:         cfg.Env.DatabasePort,
		DatabaseName: cfg.Env.DatabaseName,
	})

	services := services.NewServices(repo)

	// register api
	controller.NewHTTPController(router, services)

	http.ListenAndServe(":80", router)
}
