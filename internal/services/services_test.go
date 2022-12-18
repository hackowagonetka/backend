package services

import (
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/pkg/geocoder"
	pb_routes_analysis "backend-hagowagonetka/proto/golang/pb-routes-analysis"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newServices(t *testing.T) *Services {
	gRPCRoutesAnalysisConn, err := grpc.Dial(
		"session-manager:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	gRPCRoutesAnalysis := pb_routes_analysis.NewRoutesAnalysisClient(gRPCRoutesAnalysisConn)

	return NewServices(
		geocoder.NewYandexGeocoder("220d4c84-d54d-4a96-af30-e00235c569e3"),
		repository.NewRepository(repository.Source{
			User:         "app",
			Password:     "password",
			Host:         "localhost",
			Port:         5432,
			DatabaseName: "hackowagonetka",
		}),
		gRPCRoutesAnalysis,
	)
}
