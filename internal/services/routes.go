package services

import (
	repository_dto "backend-hagowagonetka/internal/repository/dto"
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"backend-hagowagonetka/pkg/geocoder"
	"context"
	"time"

	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

func (s *Services) RoutesHistoryDataMarshal(points repository_dto.RoutesHistoryData) ([]byte, error) {
	return json.Marshal(points)
}

func (s *Services) RoutesHistoryDataUnmarshal(raw json.RawMessage) (repository_dto.RoutesHistoryData, error) {
	var points repository_dto.RoutesHistoryData
	err := json.Unmarshal(raw, &points)
	return points, err
}

type RoutesHistoryCreatePointParam struct {
	Lon float64
	Lat float64
}

type RoutesHistoryCreateDI struct {
	UserID int32
	Points []RoutesHistoryCreatePointParam
}

func (s *Services) RoutesHistoryCreate(ctx context.Context, di RoutesHistoryCreateDI) (int64, error) {
	points := make(repository_dto.RoutesHistoryData, 0, len(di.Points))
	for index, point := range di.Points {
		geodata, err := s.Geocoder.Request(geocoder.GeocoderInput(point))
		if err != nil {
			logrus.Error(err)
			geodata.Name = "-"
		}

		points = append(points, repository_dto.RoutesHistoryPoint{
			ID:   index + 1,
			Name: geodata.Name,
			Lon:  point.Lon,
			Lat:  point.Lat,
		})
	}

	data, err := s.RoutesHistoryDataMarshal(points)
	if err != nil {
		return 0, err
	}

	return s.Repository.RoutesHistoryCreate(ctx, db_queries.RoutesHistoryCreateParams{
		CreatedAt: time.Now().UTC(),
		Data:      data,
		RefUserID: di.UserID,
	})
}

func (s *Services) RoutesHistoryGet(ctx context.Context, userID int64) ([]db_queries.RoutesHistory, error) {
	return s.Repository.RoutesHistoryGet(ctx, int32(userID))
}

type RoutesAnalysisDI struct {
	Date time.Time

	Cargo struct {
		Total  uint
		Filled uint
	}

	Points []struct {
		Lon float64 // y
		Lat float64 // x
	}
}

type RoutesAnalysisDO struct {
	Hours   uint8
	Minutes uint8
}

func (s *Services) RoutesAnalysis(ctx context.Context, di RoutesAnalysisDI) (RoutesAnalysisDO, error) {
	return RoutesAnalysisDO{}, nil
}

type RoutesDistancePoint struct {
	Lon float64
	Lat float64
}

func (s *Services) RoutesDistance(ctx context.Context, point1 RoutesDistancePoint, point2 RoutesDistancePoint) (meters float64, err error) {
	return s.Repository.RoutesDistance(ctx, db_queries.RoutesDistanceParams{
		StMakepoint:   point1.Lon,
		StMakepoint_2: point1.Lat,
		StMakepoint_3: point2.Lon,
		StMakepoint_4: point2.Lat,
	})
}
