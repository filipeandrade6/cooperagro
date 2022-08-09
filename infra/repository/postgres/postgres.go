package postgres

import (
	"context"
	"fmt"

	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO adicionar logger

type Repo struct {
	db *data.Queries
}

func NewPostgresRepo(urlConn string) (*Repo, error) {
	conn, err := pgxpool.Connect(context.Background(), urlConn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	// TODO dar o close com a aplicação

	r := data.New(conn)

	return &Repo{r}, nil
}
