package postgres

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/cooperagro/infrastructure/repository"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	DB *pgxpool.Pool
}

func NewPostgresRepo(urlConn string) (repository.Repository, error) {
	conn, err := pgxpool.Connect(context.Background(), urlConn)
	if err != nil {
		return &Repo{}, fmt.Errorf("unable to connect to database: %v", err)
	}

	// dar o close com a aplicação

	return &Repo{conn}, nil
}
