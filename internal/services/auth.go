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
