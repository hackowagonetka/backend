package services

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/repository"

	"github.com/go-chi/jwtauth/v5"
)

/*
	DI - Data Input
	DO - Data Output
*/

type Services struct {
	AuthJWT    *jwtauth.JWTAuth
	Repository *repository.Repository
}

func NewServices(
	Repository *repository.Repository,
) *Services {
	cfg := config.Get()

	return &Services{
		AuthJWT:    jwtauth.New("HS256", []byte(cfg.Env.TokenSecretKey), nil),
		Repository: Repository,
	}
}

// SELECT  CAST( ST_DistanceSphere( ST_MakePoint(50.999759,81.460766),ST_MakePoint(51.527623,81.217673) ) AS numeric );
