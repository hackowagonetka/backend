package services

import (
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/pkg/geocoder"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServices_RoutesDistance(t *testing.T) {
	services := NewServices(
		geocoder.NewYandexGeocoder("220d4c84-d54d-4a96-af30-e00235c569e3"),
		repository.NewRepository(repository.Source{
			User:         "app",
			Password:     "password",
			Host:         "localhost",
			Port:         5432,
			DatabaseName: "hackowagonetka",
		}),
	)

	meters, err := services.RouteDistance(
		context.Background(),
		RoutesDistancePoint{
			Lon: 81.460766,
			Lat: 50.999759,
		},
		RoutesDistancePoint{
			Lon: 81.217673,
			Lat: 51.527623,
		},
	)

	require.Empty(t, err)
	fmt.Println(meters)
}
