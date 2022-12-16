package app

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/controller"
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Launch() {
	cfg := config.Get()

	// create router & init base middleware
	router := chi.NewMux()
	{
		router.Use(middleware.Logger)

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           100,
		}))
	}

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
