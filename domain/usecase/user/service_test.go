package user

import (
	"testing"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		FirstName: "Filipe",
		LastName:  "Andrade",
		Address:   "Main street",
		Phone:     "5561555554444",
		Email:     "filipe@mail.com",
		Latitude:  -12.123456,
		Longitude: -12.123456,
		Roles:     []string{"admin"},
		CreatedAt: time.Now(),
	}
}

func TestCreate(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUser()
	_, err := s.CreateUser(
		u.FirstName,
		u.LastName,
		u.Address,
		u.Phone,
		u.Email,
		u.Latitude,
		u.Longitude,
		u.Roles,
	)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func TestSearchListGetUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.FirstName = "Ana"

	uID, _ := s.CreateUser(
		u1.FirstName,
		u1.LastName,
		u1.Address,
		u1.Phone,
		u1.Email,
		u1.Latitude,
		u1.Longitude,
		u1.Roles,
	)
	_, _ = s.CreateUser(
		u2.FirstName,
		u2.LastName,
		u2.Address,
		u2.Phone,
		u2.Email,
		u2.Latitude,
		u2.Longitude,
		u2.Roles,
	)

	t.Run("search", func(t *testing.T) {
		u, err := s.SearchUser("ANA")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(u))
		assert.Equal(t, "Ana", u[0].FirstName)

		u, err = s.SearchUser("josé")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, u)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := s.ListUser()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		u, err := s.GetUserByID(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.FirstName, u.FirstName)
	})
}

// TODO teste update para nome existente (que deveria ser unico)

func TestUpdateDeleteUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUser()

	id, err := s.CreateUser(
		u.FirstName,
		u.LastName,
		u.Address,
		u.Phone,
		u.Email,
		u.Latitude,
		u.Longitude,
		u.Roles,
	)
	assert.Nil(t, err)
	saved, _ := s.GetUserByID(id)
	saved.FirstName = "ana"
	assert.Nil(t, s.UpdateUser(saved))
	updated, err := s.GetUserByID(id)
	assert.Nil(t, err)
	assert.Equal(t, "ana", updated.FirstName)

	assert.Nil(t, s.DeleteUser(id))
	assert.Equal(t, entity.ErrNotFound, s.DeleteUser(id))
}
