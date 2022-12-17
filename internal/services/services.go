package services

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/pkg/geocoder"

	"github.com/go-chi/jwtauth/v5"
)

/*
	DI - Data Input
	DO - Data Output
*/

type Services struct {
	AuthJWT    *jwtauth.JWTAuth
	Geocoder   geocoder.Geocoder
	Repository *repository.Repository
}

func NewServices(
	Geocoder geocoder.Geocoder,
	Repository *repository.Repository,
) *Services {
	cfg := config.Get()

	return &Services{
		AuthJWT:    jwtauth.New("HS256", []byte(cfg.Env.TokenSecretKey), nil),
		Geocoder:   Geocoder,
		Repository: Repository,
	}
}
