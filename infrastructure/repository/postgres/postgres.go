package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	DB *pgxpool.Pool
}

func NewPostgresRepo(urlConn string) (*Repository, error) {
	conn, err := pgxpool.Connect(context.Background(), urlConn)
	if err != nil {
		return &Repos{}, fmt.Errorf("unable to connect to database: %v", err)
	}

	// dar o close com a aplicação

	return &Repos{conn}, nil
}
