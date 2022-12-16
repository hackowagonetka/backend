package services

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/repository"
	"fmt"

	"github.com/go-chi/jwtauth/v5"
)

type Services struct {
	jwtAuth    *jwtauth.JWTAuth
	Repository *repository.Repository
}

func NewServices(
	Repository *repository.Repository,
) *Services {

	cfg := config.Get()

	fmt.Println(cfg.Env.TokenSecretKey)

	return &Services{
		jwtAuth:    jwtauth.New("HS256", []byte(cfg.Env.TokenSecretKey), nil),
		Repository: Repository,
	}
}
