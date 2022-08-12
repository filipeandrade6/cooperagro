package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO adicionar logger

type Repo struct {
	db *data.Queries
}

type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	TLS      string
}

func NewPostgresRepo(cfg Config) (*Repo, error) {
	q := make(url.Values)
	q.Set("sslmode", cfg.TLS)

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	conn, err := pgxpool.Connect(context.Background(), u.String())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	// TODO dar o close com a aplicação

	r := data.New(conn)

	return &Repo{r}, nil
}
