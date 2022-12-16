package controller

import "github.com/go-chi/chi/v5"

func (c *HTTPController) register(router *chi.Mux) {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", c.AuthSignUp)
	})
}
