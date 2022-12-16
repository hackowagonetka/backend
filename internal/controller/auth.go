package controller

import (
	"backend-hagowagonetka/internal/controller/render"
	"backend-hagowagonetka/internal/services"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

	id, err := c.Services.AuthSignUp(r.Context(), services.AuthSignUpDI{
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

		logrus.Error(fmt.Errorf("controller: AuthSignUp: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	render.NewReponse(http.StatusCreated, w, AuthSignUpResponse{
		ID: id,
	})
}

type AuthSignInBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthSignInResponse struct {
	Token string `json:"token"`
}

func (c *HTTPController) AuthSignIn(w http.ResponseWriter, r *http.Request) {
	var body AuthSignInBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		render.NewReponse(http.StatusBadRequest, w, err)
		return
	}

	token, err := c.Services.AuthSignIn(r.Context(), services.AuthSignInDI{
		Login:    body.Login,
		Password: body.Password,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.NewReponse(http.StatusNotFound, w, errors.New("login not found"))
			return
		} else if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			render.NewReponse(http.StatusUnauthorized, w, errors.New("invalid password"))
			return
		}

		logrus.Error(fmt.Errorf("controller: AuthSignUp: %w", err))
		render.NewReponse(http.StatusInternalServerError, w, nil)
		return
	}

	render.NewReponse(http.StatusOK, w, AuthSignInResponse{
		Token: token,
	})
}
