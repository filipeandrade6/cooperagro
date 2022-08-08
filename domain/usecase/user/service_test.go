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
	u2.Name = "manga"

	uID, _ := s.CreateUser(u1.Name)
	_, _ = s.CreateUser(u2.Name)

	t.Run("search", func(t *testing.T) {
		u, err := s.SearchUser("MANGA")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(u))
		assert.Equal(t, "manga", u[0].Name)

		u, err = s.SearchUser("acerola")
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
		assert.Equal(t, u1.Name, u.Name)
	})
}

// TODO teste update para nome existente (que deveria ser unico)

func TestUpdateDeleteUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUser()

	id, err := s.CreateUser(u.Name)
	assert.Nil(t, err)
	saved, _ := s.GetUserByID(id)
	saved.Name = "manga"
	assert.Nil(t, s.UpdateUser(saved))
	updated, err := s.GetUserByID(id)
	assert.Nil(t, err)
	assert.Equal(t, "manga", updated.Name)

	assert.Nil(t, s.DeleteUser(id))
	assert.Equal(t, entity.ErrNotFound, s.DeleteUser(id))
}
