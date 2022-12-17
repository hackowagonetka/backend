package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServices_RoutesConverTime(t *testing.T) {
	services := newServices(t)
	tm := services.RoutesTimeConvert(1506)

	fmt.Println("days: ", tm.Days)
	fmt.Println("hours: ", tm.Hours)
	fmt.Println("minutes: ", tm.Minutes)
}

func TestServices_RoutesDistance(t *testing.T) {
	services := newServices(t)

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
