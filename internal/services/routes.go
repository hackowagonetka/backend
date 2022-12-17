package services

import (
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"context"
	"time"
)

type RoutesAnalysisDI struct {
	Date  time.Time
	Cargo struct {
		Total  uint
		Filled uint
	}
	Stations []int64
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

func (s *Services) RouteDistance(ctx context.Context, p1 RoutesDistancePoint, p2 RoutesDistancePoint) (meters float64, err error) {
	return s.Repository.RouteDistance(ctx, db_queries.RouteDistanceParams{
		StMakepoint:   p1.Lon,
		StMakepoint_2: p1.Lat,
		StMakepoint_3: p2.Lon,
		StMakepoint_4: p2.Lat,
	})
}
