package services

import (
	"backend-hagowagonetka/internal/repository"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServices_RoutesDistance(t *testing.T) {
	services := NewServices(repository.NewRepository(repository.Source{
		User:         "app",
		Password:     "password",
		Host:         "localhost",
		Port:         5432,
		DatabaseName: "hackowagonetka",
	}))

	meters, err := services.RoutesDistance(
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
