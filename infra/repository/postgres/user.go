package postgres

import (
	"context"
	"errors"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/infra/repository/postgres/data"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (r *Repo) GetUserByID(id entity.ID) (*entity.User, error) {
	ctx := context.Background()
	c, err := r.db.GetUserByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		Phone:     c.Phone,
		Email:     c.Email,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
		Roles:     c.Roles,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (r *Repo) SearchUser(query string) ([]*entity.User, error) {
	ctx := context.Background()
	users, err := r.db.SearchUser(ctx, query) // TODO SearchUser faz busca na coluna first_name -> alterar depois
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var usersOut []*entity.User
	for _, user := range users {
		usersOut = append(usersOut, &entity.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Address:   user.Address,
			Phone:     user.Phone,
			Email:     user.Email,
			Latitude:  user.Latitude,
			Longitude: user.Longitude,
			Roles:     user.Roles,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return usersOut, nil
}

func (r *Repo) ListUser() ([]*entity.User, error) {
	ctx := context.Background()
	users, err := r.db.ListUser(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	var usersOut []*entity.User
	for _, User := range users {
		usersOut = append(usersOut, &entity.User{
			ID:        User.ID,
			FirstName: User.FirstName,
			LastName:  User.LastName,
			Address:   User.Address,
			Phone:     User.Phone,
			Email:     User.Email,
			Latitude:  User.Latitude,
			Longitude: User.Longitude,
			Roles:     User.Roles,
			CreatedAt: User.CreatedAt,
			UpdatedAt: User.UpdatedAt,
		})
	}

	return usersOut, nil
}

func (r *Repo) CreateUser(e *entity.User) (entity.ID, error) {
	ctx := context.Background()
	_, err := r.db.CreateUser(ctx, data.CreateUserParams{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Address:   e.Address,
		Phone:     e.Phone,
		Email:     e.Email,
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Roles:     e.Roles,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	})

	if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code == "23505" {
		return e.ID, entity.ErrEntityAlreadyExists
	}

	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

func (r *Repo) UpdateUser(e *entity.User) error {
	ctx := context.Background()
	err := r.db.UpdateUser(ctx, data.UpdateUserParams{
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Address:   e.Address,
		Phone:     e.Phone,
		Email:     e.Email,
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Roles:     e.Roles,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		ID:        e.ID,
	})

	if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code == "23505" {
		return entity.ErrEntityAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteUser(id entity.ID) error {
	ctx := context.Background()
	err := r.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
