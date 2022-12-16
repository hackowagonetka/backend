package controller

import (
	"backend-hagowagonetka/internal/controller/render"
	"backend-hagowagonetka/internal/services"
	"context"
	"errors"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/lib/pq"
)

type AuthSignUpBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthSignUpResponse struct {
	ID int64 `json:"id"`
}

func (c *HTTPController) AuthSignUp(w http.ResponseWriter, r *http.Request) {
	var body AuthSignUpBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.NewReponse(http.StatusBadRequest, w, err)
		return
	}

	id, err := c.Services.AuthSignUp(context.Background(), services.AuthSignUpDI{
		Login:    body.Login,
		Password: body.Password,
	})
	if err != nil {
		if dberr, ok := err.(*pq.Error); ok {
			switch dberr.Code.Name() {
			case "unique_violation":
				render.NewReponse(http.StatusForbidden, w, errors.New("login already exists"))
				return
			}
		}

		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	render.NewReponse(http.StatusCreated, w, AuthSignUpResponse{
		ID: id,
	})
}
