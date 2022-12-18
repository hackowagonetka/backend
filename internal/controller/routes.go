package controller

import (
	"backend-hagowagonetka/internal/controller/render"
	"backend-hagowagonetka/internal/services"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type RoutesAnalysisBody struct {
	Date  time.Time `json:"date"`
	Cargo struct {
		Total  int32 `json:"total"`
		Filled int32 `json:"filled"`
	} `json:"cargo"`
	Stations []int64 `json:"stations"`
}

type RoutesAnalysisResponse struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

func (c *HTTPController) RouteAnalysis(w http.ResponseWriter, r *http.Request) {
	var body RoutesAnalysisBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.NewReponse(http.StatusBadRequest, w, err)
		return
	}

	analys, err := c.Services.RoutesAnalysis(r.Context(), services.RoutesAnalysisDI(body))
	if err != nil {
		logrus.Error(fmt.Errorf("controller: RoutesAnalysis: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	t := c.Services.RoutesTimeConvert(analys.TimeSpent)
	render.NewReponse(http.StatusOK, w, RoutesAnalysisResponse(t))
}
