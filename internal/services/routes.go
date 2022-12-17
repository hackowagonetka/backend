package services

import (
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"context"
	"time"
)

const MINUTES_PER_DAY = 1440

type RoutesAnalysisTimeConvert struct {
	Days    int
	Hours   int
	Minutes int
}

func (s *Services) RoutesTimeConvert(totalMinutes int64) RoutesAnalysisTimeConvert {
	return RoutesAnalysisTimeConvert{
		Days:    int(totalMinutes / MINUTES_PER_DAY),
		Hours:   int((totalMinutes % MINUTES_PER_DAY) / 60),
		Minutes: int((totalMinutes % MINUTES_PER_DAY) % 60),
	}
}

type RoutesAnalysisDI struct {
	Date  time.Time
	Cargo struct {
		Total  uint
		Filled uint
	}
	Stations []int64
}

type RoutesAnalysisDO struct {
	TimeSpent int64 // in minutes
}

func (s *Services) RoutesAnalysis(ctx context.Context, di RoutesAnalysisDI) (RoutesAnalysisDO, error) {
	return RoutesAnalysisDO{}, nil
}

type RoutesDistancePoint struct {
	Lon float64
	Lat float64
}

func (s *Services) RoutesDistance(ctx context.Context, p1 RoutesDistancePoint, p2 RoutesDistancePoint) (meters float64, err error) {
	return s.Repository.RouteDistance(ctx, db_queries.RouteDistanceParams{
		StMakepoint:   p1.Lon,
		StMakepoint_2: p1.Lat,
		StMakepoint_3: p2.Lon,
		StMakepoint_4: p2.Lat,
	})
}
