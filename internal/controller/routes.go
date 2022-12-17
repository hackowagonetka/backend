package controller

import (
	"backend-hagowagonetka/internal/controller/render"
	repository_dto "backend-hagowagonetka/internal/repository/dto"
	"backend-hagowagonetka/internal/services"
	"backend-hagowagonetka/pkg/geocoder"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type RoutesAnalysisBody struct {
	Date  time.Time `json:"date"`
	Cargo struct {
		Total  uint `json:"total"`
		Filled uint `json:"filled"`
	} `json:"cargo"`
	Points []struct {
		Lon float64 `json:"lng"` // y
		Lat float64 `json:"lat"` // x
	} `json:"points"`
}

type RoutesAnalysisResponse struct {
	Hours   uint8 `json:"hours"`
	Minutes uint8 `json:"minutes"`
}

func (c *HTTPController) RoutesAnalysis(w http.ResponseWriter, r *http.Request) {
	var body RoutesAnalysisBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.NewReponse(http.StatusBadRequest, w, err)
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err == nil {
		user_id := int64(claims["user_id"].(float64))

		points := make(repository_dto.RoutesHistoryData, 0, len(body.Points))
		for index, point := range body.Points {
			geodata, err := c.Services.Geocoder.Request(geocoder.GeocoderInput(point))
			if err != nil {
				logrus.Error(fmt.Errorf("controller: RoutesAnalysis: %w", err))
				render.NewReponse(http.StatusInternalServerError, w, nil)
			}

			points = append(points, repository_dto.RoutesHistoryPoint{
				ID:   index + 1,
				Name: geodata.Name,
				Lon:  point.Lon,
				Lat:  point.Lat,
			})
		}

		// save history
		if _, err := c.Services.RoutesHistoryCreate(r.Context(), services.RoutesHistoryCreateDI{
			UserID: int32(user_id),
			Points: points,
		}); err != nil {
			logrus.Error(fmt.Errorf("controller: RoutesAnalysis: %w", err))
			render.NewReponse(http.StatusInternalServerError, w, nil)
			return
		}
	}

	analys, err := c.Services.RoutesAnalysis(r.Context(), services.RoutesAnalysisDI(body))
	if err != nil {
		logrus.Error(fmt.Errorf("controller: RoutesAnalysis: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	render.NewReponse(http.StatusOK, w, RoutesAnalysisResponse(analys))
}

type RoutesHistoryGetData struct {
	ID        int64                            `json:"id"`
	CreatedAt time.Time                        `json:"created_at"`
	Points    repository_dto.RoutesHistoryData `json:"points"`
}

type RoutesHistoryGetResponse []RoutesHistoryGetData

func (c *HTTPController) RoutesHistoryGet(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.NewReponse(http.StatusBadRequest, w, err)
		return
	}

	user_id := int64(claims["user_id"].(float64))

	histories, err := c.Services.RoutesHistoryGet(r.Context(), user_id)
	if err != nil {
		logrus.Error(fmt.Errorf("controller: RoutesHistoryGet: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	var response RoutesHistoryGetResponse = make(RoutesHistoryGetResponse, 0, len(histories))
	for _, history := range histories {
		points, err := c.Services.RoutesUnmarshalPoints(history.Data)
		if err != nil {
			logrus.Error(fmt.Errorf("controller: RoutesHistoryGet: %w", err))
			render.NewReponse(http.StatusInternalServerError, w, nil)
			return
		}

		response = append(response, RoutesHistoryGetData{
			ID:        history.ID,
			CreatedAt: history.CreatedAt,
			Points:    points,
		})
	}

	render.NewReponse(http.StatusOK, w, response)
}