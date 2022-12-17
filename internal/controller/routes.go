package controller

import (
	"backend-hagowagonetka/internal/controller/render"
	"backend-hagowagonetka/internal/services"
	"fmt"
	"net/http"
	"time"

	"github.com/cridenour/go-postgis"
	"github.com/go-chi/jwtauth/v5"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type RoutesAnalysisBody struct {
	Date time.Time `json:"time"`

	Cargo struct {
		Total  uint `json:"total"`
		Filled uint `json:"filled"`
	} `json:"cargo"`

	Points []struct {
		Lon float64 `json:"lon"` // y
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
	if err != nil {
		user_id := claims["user_id"].(int64)
		points := make([]postgis.Point, 0, len(body.Points))

		for _, point := range body.Points {
			points = append(points, postgis.Point{
				X: point.Lat,
				Y: point.Lon,
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
