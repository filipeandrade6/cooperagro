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

func TestService_GetUserByID(t *testing.T) {
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

	t.Run("get existent", func(t *testing.T) {
		saved, err := s.GetUserByID(id)
		assert.Nil(t, err)
		assert.Equal(t, u.FirstName, saved.FirstName)
		assert.Equal(t, u.LastName, saved.LastName)
		assert.Equal(t, u.Address, saved.Address)
		assert.Equal(t, u.Phone, saved.Phone)
		assert.Equal(t, u.Email, saved.Email)
		assert.Equal(t, u.Latitude, saved.Latitude)
		assert.Equal(t, u.Longitude, saved.Longitude)
		assert.Equal(t, u.Roles, saved.Roles)
	})

	t.Run("get non existent", func(t *testing.T) {
		_, err := s.GetUserByID(entity.NewID())
		assert.Equal(t, entity.ErrNotFound, err)
	})
}

func TestService_SearchUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUser()

	_, _ = s.CreateUser(
		u.FirstName,
		u.LastName,
		u.Address,
		u.Phone,
		u.Email,
		u.Latitude,
		u.Longitude,
		u.Roles,
	)

	t.Run("search equal", func(t *testing.T) {
		saved, err := s.SearchUser("Filipe")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, u.FirstName, saved[0].FirstName)
		assert.Equal(t, u.LastName, saved[0].LastName)
		assert.Equal(t, u.Address, saved[0].Address)
		assert.Equal(t, u.Phone, saved[0].Phone)
		assert.Equal(t, u.Email, saved[0].Email)
		assert.Equal(t, u.Latitude, saved[0].Latitude)
		assert.Equal(t, u.Longitude, saved[0].Longitude)
		assert.Equal(t, u.Roles, saved[0].Roles)
	})

	t.Run("search equal but with capital letters", func(t *testing.T) {
		saved, err := s.SearchUser("FILIPE")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, u.FirstName, saved[0].FirstName)
		assert.Equal(t, u.LastName, saved[0].LastName)
		assert.Equal(t, u.Address, saved[0].Address)
		assert.Equal(t, u.Phone, saved[0].Phone)
		assert.Equal(t, u.Email, saved[0].Email)
		assert.Equal(t, u.Latitude, saved[0].Latitude)
		assert.Equal(t, u.Longitude, saved[0].Longitude)
		assert.Equal(t, u.Roles, saved[0].Roles)
	})

	t.Run("search for inexistent", func(t *testing.T) {
		saved, err := s.SearchUser("João")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(saved))
	})
}

func TestService_ListUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)

	t.Run("list empty", func(t *testing.T) {
		us, err := s.ListUser()
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(us))
	})

	t.Run("list all", func(t *testing.T) {
		u1 := newFixtureUser()
		u2 := newFixtureUser()
		u2.Email = "other@email.com"
		_, _ = s.CreateUser(
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

		us, err := s.ListUser()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(us))
	})
}

func TestService_CreateUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUser()

	t.Run("create user", func(t *testing.T) {
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
	})

	t.Run("create existent user", func(t *testing.T) {
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
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_UpdateUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)

	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.Email = "other@email.com"

	id, _ := s.CreateUser(
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

	t.Run("update user", func(t *testing.T) {
		err := s.UpdateUser(&entity.User{
			ID:        id,
			FirstName: "Marcelo",
			LastName:  u1.LastName,
			Address:   u1.Address,
			Phone:     u1.Phone,
			Email:     u1.Email,
			Latitude:  u1.Latitude,
			Longitude: u1.Longitude,
			Roles:     u1.Roles,
		})
		assert.Nil(t, err)
	})

	t.Run("update to existent user", func(t *testing.T) {
		err := s.UpdateUser(&entity.User{
			ID:        id,
			FirstName: "Marcelo",
			LastName:  u1.LastName,
			Address:   u1.Address,
			Phone:     u1.Phone,
			Email:     "other@email.com",
			Latitude:  u1.Latitude,
			Longitude: u1.Longitude,
			Roles:     u1.Roles,
		})
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_DeleteUser(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUser()
	id, _ := s.CreateUser(
		u.FirstName,
		u.LastName,
		u.Address,
		u.Phone,
		u.Email,
		u.Latitude,
		u.Longitude,
		u.Roles,
	)

	t.Run("delete user", func(t *testing.T) {
		err := s.DeleteUser(id)
		assert.Nil(t, err)
	})

	t.Run("delete inexistent user", func(t *testing.T) {
		err := s.DeleteUser(id)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}
