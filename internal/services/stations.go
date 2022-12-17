package services

import (
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"backend-hagowagonetka/pkg/geocoder"
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type StationCreateDI struct {
	Name   string
	Lon    float64
	Lat    float64
	UserID int64
}

func (s *Services) StationCreate(ctx context.Context, di StationCreateDI) (int64, error) {
	params := db_queries.StationCreateParams{
		CreatedAt: time.Now().UTC(),
		Name:      di.Name,
		Lon:       di.Lon,
		Lat:       di.Lat,
		RefUserID: di.UserID,
	}

	geodata, err := s.Geocoder.Request(geocoder.GeocoderInput{
		Lon: di.Lon,
		Lat: di.Lat,
	})
	if err != nil {
		logrus.Error(err)
	}
	params.Geoname = geodata.Name

	return s.Repository.StationCreate(ctx, params)
}

func (s *Services) StationGetList(ctx context.Context, userID int64) ([]db_queries.Station, error) {
	return s.Repository.StationGetList(ctx, userID)
}
