package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func (c *HTTPController) register(router *chi.Mux) {
	router.Use(jwtauth.Verifier(c.Services.AuthJWT))

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", c.AuthSignUp)
		r.Post("/sign-in", c.AuthSignIn)
	})

	router.Route("/routes", func(r chi.Router) {
		r.Post("/", c.RoutesAnalysis)
		r.With(jwtauth.Authenticator).Get("/", c.RoutesHistoryGet)
	})
}
