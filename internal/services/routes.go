package services

import (
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	pb_routes_analysis "backend-hagowagonetka/proto/golang/pb-routes-analysis"
	"context"
	"errors"
	"time"
)

const MINUTES_PER_DAY = 1440

var (
	ErrMinimumNumberOfStations = errors.New("minimum number of stations: 2")
)

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
		Total  int32
		Filled int32
	}
	Stations []int64
}

type RoutesAnalysisDO struct {
	TimeSpent int64 // in minutes
}

func (s *Services) RoutesAnalysis(ctx context.Context, di RoutesAnalysisDI) (RoutesAnalysisDO, error) {
	stations, err := s.Repository.StationGetListByID(ctx, di.Stations)
	if err != nil {
		return RoutesAnalysisDO{}, err
	}

	if len(stations) < 2 {
		return RoutesAnalysisDO{}, ErrMinimumNumberOfStations
	}

	distanceM, err := s.RoutesDistance(ctx,
		RoutesDistancePoint{
			Lon: stations[0].Lon,
			Lat: stations[0].Lat,
		},
		RoutesDistancePoint{
			Lon: stations[0].Lon,
			Lat: stations[0].Lat,
		},
	)
	if err != nil {
		return RoutesAnalysisDO{}, err
	}

	distanceKM := distanceM / 1000

	analyse, err := s.gRPCRoutesAnalysis.Analyse(ctx, &pb_routes_analysis.AnalyseRequest{
		Distance:    int64(distanceKM),
		Timestamp:   di.Date.Unix(),
		CargoTotal:  int32(di.Cargo.Total),
		CargoFilled: int32(di.Cargo.Filled),
	})
	if err != nil {
		return RoutesAnalysisDO{}, err
	}

	return RoutesAnalysisDO{
		TimeSpent: analyse.TimeSpent * 60,
	}, nil
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
