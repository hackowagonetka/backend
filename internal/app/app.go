package app

import (
	"backend-hagowagonetka/internal/config"
	"backend-hagowagonetka/internal/controller"
	"backend-hagowagonetka/internal/repository"
	"backend-hagowagonetka/internal/services"
	"backend-hagowagonetka/pkg/geocoder"
	pb_routes_analysis "backend-hagowagonetka/proto/golang/pb-routes-analysis"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Launch() {
	cfg := config.Get()

	// create router & init base middleware
	router := chi.NewMux()
	{
		router.Use(middleware.Logger)

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           100,
		}))
	}

	geocoder := geocoder.NewYandexGeocoder(cfg.Env.YandexGeocoderToken)

	repo := repository.NewRepository(repository.Source{
		User:         cfg.Env.DatabaseUser,
		Password:     cfg.Env.DatabasePassword,
		Host:         cfg.Env.DatabaseHost,
		Port:         cfg.Env.DatabasePort,
		DatabaseName: cfg.Env.DatabaseName,
	})

	//connect to gRPC Routes Analysis
	gRPCRoutesAnalysisConn, err := grpc.Dial(
		"localhost:7878",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer gRPCRoutesAnalysisConn.Close()
	gRPCRoutesAnalysis := pb_routes_analysis.NewRoutesAnalysisClient(gRPCRoutesAnalysisConn)

	services := services.NewServices(geocoder, repo, gRPCRoutesAnalysis)

	// register api
	controller.NewHTTPController(router, services)

	http.ListenAndServe(":80", router)
}
