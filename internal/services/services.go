package services

import "backend-hagowagonetka/internal/repository"

type Services struct {
	Repository *repository.Repository
}

func NewServices(
	Repository *repository.Repository,
) *Services {
	return &Services{
		Repository: Repository,
	}
}
