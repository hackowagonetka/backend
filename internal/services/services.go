package services

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/pkg/geocoder"
	pb_routes_analysis "backend-hagowagonetka/proto/golang/pb-routes-analysis"

	"github.com/go-chi/jwtauth/v5"
)

/*
	DI - Data Input
	DO - Data Output
*/

type Services struct {
	AuthJWT            *jwtauth.JWTAuth
	Geocoder           geocoder.Geocoder
	Repository         *repository.Repository
	gRPCRoutesAnalysis pb_routes_analysis.RoutesAnalysisClient
}

func NewServices(
	Geocoder geocoder.Geocoder,
	Repository *repository.Repository,
	gRPCRoutesAnalysis pb_routes_analysis.RoutesAnalysisClient,
) *Services {
	cfg := config.Get()

	return &Services{
		AuthJWT:            jwtauth.New("HS256", []byte(cfg.Env.TokenSecretKey), nil),
		Geocoder:           Geocoder,
		Repository:         Repository,
		gRPCRoutesAnalysis: gRPCRoutesAnalysis,
	}
}
