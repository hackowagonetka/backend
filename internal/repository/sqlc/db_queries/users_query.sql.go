// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users_query.sql

package db_queries

import (
	"context"
)

const userCreate = `-- name: UserCreate :one
INSERT INTO users (
    login, password
) VALUES (
    $1, $2
) RETURNING id
`

type UserCreateParams struct {
	Login    string
	Password string
}

func (q *Queries) UserCreate(ctx context.Context, arg UserCreateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, userCreate, arg.Login, arg.Password)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const userGet = `-- name: UserGet :one
SELECT id, login, password FROM users WHERE login = $1
`

func (q *Queries) UserGet(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, userGet, login)
	var i User
	err := row.Scan(&i.ID, &i.Login, &i.Password)
	return i, err
}
