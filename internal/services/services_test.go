package services

import (
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/pkg/geocoder"
	"testing"
)

func newServices(t *testing.T) *Services {
	return NewServices(
		geocoder.NewYandexGeocoder("220d4c84-d54d-4a96-af30-e00235c569e3"),
		repository.NewRepository(repository.Source{
			User:         "app",
			Password:     "password",
			Host:         "localhost",
			Port:         5432,
			DatabaseName: "hackowagonetka",
		}),
	)
}
