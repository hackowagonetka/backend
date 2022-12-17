package services

import (
	"backend-hagowagonetka/internal/repository/dto"
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"context"
	"time"

	"github.com/cridenour/go-postgis"
)

type RoutesHistoryCreateDI struct {
	UserID int32
	Points []postgis.Point
}

func (s *Services) RoutesHistoryCreate(ctx context.Context, di RoutesHistoryCreateDI) (int64, error) {
	return s.Repository.RoutesHistoryCreate(ctx, db_queries.RoutesHistoryCreateParams{
		CreatedAt: time.Now().UTC(),
		Data: &dto.RoutesHistoryData{
			Points: di.Points,
		},
		RefUserID: di.UserID,
	})
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
