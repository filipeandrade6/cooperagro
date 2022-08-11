package unitofmeasure

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureUnitOfMeasure() *entity.UnitOfMeasure {
	return &entity.UnitOfMeasure{
		Name: "kilogram",
	}
}

func TestService_GetUnitOfMeasureByID(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUnitOfMeasure()
	id, err := s.CreateUnitOfMeasure(u.Name)
	assert.Nil(t, err)

	t.Run("get existent", func(t *testing.T) {
		saved, err := s.GetUnitOfMeasureByID(id)
		assert.Nil(t, err)
		assert.Equal(t, u.Name, saved.Name)
	})

	t.Run("get non existent", func(t *testing.T) {
		_, err := s.GetUnitOfMeasureByID(entity.NewID())
		assert.Equal(t, entity.ErrNotFound, err)
	})
}

func TestService_SearchUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUnitOfMeasure()

	_, _ = s.CreateUnitOfMeasure(u.Name)

	t.Run("search equal", func(t *testing.T) {
		saved, err := s.SearchUnitOfMeasure("kilogram")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, u.Name, saved[0].Name)
	})

	t.Run("search equal but with capital letters", func(t *testing.T) {
		saved, err := s.SearchUnitOfMeasure("KilograM")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, u.Name, saved[0].Name)
	})

	t.Run("search for inexistent", func(t *testing.T) {
		saved, err := s.SearchUnitOfMeasure("piece")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(saved))
	})
}

func TestService_ListUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)

	t.Run("list empty", func(t *testing.T) {
		us, err := s.ListUnitOfMeasure()
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(us))
	})

	t.Run("list all", func(t *testing.T) {
		u1 := newFixtureUnitOfMeasure()
		u2 := newFixtureUnitOfMeasure()
		u2.Name = "piece"
		_, _ = s.CreateUnitOfMeasure(u1.Name)
		_, _ = s.CreateUnitOfMeasure(u2.Name)

		us, err := s.ListUnitOfMeasure()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(us))
	})
}

func TestService_CreateUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUnitOfMeasure()

	t.Run("create unit of measure", func(t *testing.T) {
		_, err := s.CreateUnitOfMeasure(u.Name)
		assert.Nil(t, err)
	})

	t.Run("create existent unit of measure", func(t *testing.T) {
		_, err := s.CreateUnitOfMeasure(u.Name)
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_UpdateUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)

	u := newFixtureUnitOfMeasure()
	u2 := newFixtureUnitOfMeasure()
	u2.Name = "piece"

	id, _ := s.CreateUnitOfMeasure(u.Name)
	_, _ = s.CreateUnitOfMeasure(u2.Name)

	t.Run("update unit of measure", func(t *testing.T) {
		err := s.UpdateUnitOfMeasure(&entity.UnitOfMeasure{
			ID:   id,
			Name: "liters",
		})
		assert.Nil(t, err)
	})

	t.Run("update to existent unit of measure", func(t *testing.T) {
		err := s.UpdateUnitOfMeasure(&entity.UnitOfMeasure{
			ID:   id,
			Name: u2.Name,
		})
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_DeleteUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureUnitOfMeasure()
	id, _ := s.CreateUnitOfMeasure(u.Name)

	t.Run("delete unit of measure", func(t *testing.T) {
		err := s.DeleteUnitOfMeasure(id)
		assert.Nil(t, err)
	})

	t.Run("delete inexistent unit of measure", func(t *testing.T) {
		err := s.DeleteUnitOfMeasure(id)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}
