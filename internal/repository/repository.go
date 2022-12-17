package repository

import (
	"backend-hagowagonetka/internal/repository/sqlc/db_queries"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Source struct {
	User         string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

type Repository struct {
	db *sql.DB
	*db_queries.Queries
}

func NewRepository(src Source) *Repository {
	link := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", src.User, src.Password, src.Host, src.Port, src.DatabaseName)

	db, err := sql.Open("postgres", link)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Repository{
		db:      db,
		Queries: db_queries.New(db),
	}
}

func (r *Repository) TX() (*sql.Tx, error) {
	return r.db.Begin()
}
