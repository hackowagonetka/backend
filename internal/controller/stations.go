package controller

import (
	"backend-hagowagonetka/internal/controller/render"
	"backend-hagowagonetka/internal/services"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type StationCreateBody struct {
	Name string  `json:"name"`
	Lon  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

type StationCreateReponse struct {
	ID int64 `json:"id"`
}

func (s *HTTPController) StationCreate(w http.ResponseWriter, r *http.Request) {
	var body StationCreateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.NewReponse(http.StatusBadRequest, w, err)
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.NewReponse(http.StatusInternalServerError, w, err)
		return
	}

	userID := int64(claims["user_id"].(float64))

	id, err := s.Services.StationCreate(r.Context(), services.StationCreateDI{
		Name:   body.Name,
		Lon:    body.Lon,
		Lat:    body.Lat,
		UserID: userID,
	})
	if err != nil {
		logrus.Error(fmt.Errorf("controller: StationCreate: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	render.NewReponse(http.StatusCreated, w, StationCreateReponse{
		ID: id,
	})
}

type StationGetListItem struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Geoname   string    `json:"geoname"`
	Lon       float64   `json:"lng"`
	Lat       float64   `json:"lat"`
}

type StationGetListResponse []StationGetListItem

func (s *HTTPController) StationGetList(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.NewReponse(http.StatusInternalServerError, w, err)
		return
	}

	userID := int64(claims["user_id"].(float64))

	list, err := s.Services.StationGetList(r.Context(), userID)
	if err != nil {
		logrus.Error(fmt.Errorf("controller: StationGetList: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	response := make(StationGetListResponse, 0, len(list))
	for _, station := range list {
		response = append(response, StationGetListItem{
			ID:        station.ID,
			CreatedAt: station.CreatedAt,
			Name:      station.Name,
			Geoname:   station.Geoname,
			Lon:       station.Lon,
			Lat:       station.Lat,
		})
	}

	render.NewReponse(http.StatusOK, w, response)
}
