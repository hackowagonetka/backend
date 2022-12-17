package services

import (
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type AuthSignUpDI struct {
	Login    string
	Password string
}

func (s *Services) AuthSignUp(ctx context.Context, di AuthSignUpDI) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(di.Password), 10)
	if err != nil {
		return 0, err
	}

	return s.Repository.UserCreate(ctx, db_queries.UserCreateParams{
		Login:    di.Login,
		Password: string(hash),
	})
}

type AuthSignInDI struct {
	Login    string
	Password string
}

func (s *Services) AuthSignIn(ctx context.Context, di AuthSignInDI) (string, error) {
	user, err := s.Repository.UserGet(ctx, di.Login)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(di.Password)); err != nil {
		return "", err
	}

	_, token, err := s.AuthJWT.Encode(map[string]interface{}{
		"user_id": user.ID,
	})

	return token, err
}
