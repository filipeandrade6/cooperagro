package unitofmeasure

import (
	"testing"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureUnitOfMeasure() *entity.UnitOfMeasure {
	return &entity.UnitOfMeasure{
		Name:      "kilogram",
		CreatedAt: time.Now(),
	}
}

func TestCreate(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureUnitOfMeasure()
	_, err := s.CreateUnitOfMeasure((bp.Name))
	assert.Nil(t, err)
	assert.False(t, bp.CreatedAt.IsZero())
}

func TestSearchListGetUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp1 := newFixtureUnitOfMeasure()
	bp2 := newFixtureUnitOfMeasure()
	bp2.Name = "unit"

	uID, _ := s.CreateUnitOfMeasure(bp1.Name)
	_, _ = s.CreateUnitOfMeasure(bp2.Name)

	t.Run("search", func(t *testing.T) {
		bp, err := s.SearchUnitOfMeasure("UNIT")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(bp))
		assert.Equal(t, "unit", bp[0].Name)

		bp, err = s.SearchUnitOfMeasure("liters")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, bp)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := s.ListUnitOfMeasure()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		bp, err := s.GetUnitOfMeasureByID(uID)
		assert.Nil(t, err)
		assert.Equal(t, bp1.Name, bp.Name)
	})
}

// TODO teste update para nome existente (que deveria ser unico)

func TestUpdateDeleteUnitOfMeasure(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureUnitOfMeasure()

	id, err := s.CreateUnitOfMeasure(bp.Name)
	assert.Nil(t, err)
	saved, _ := s.GetUnitOfMeasureByID(id)
	saved.Name = "unit"
	assert.Nil(t, s.UpdateUnitOfMeasure(saved))
	updated, err := s.GetUnitOfMeasureByID(id)
	assert.Nil(t, err)
	assert.Equal(t, "unit", updated.Name)

	assert.Nil(t, s.DeleteUnitOfMeasure(id))
	assert.Equal(t, entity.ErrNotFound, s.DeleteUnitOfMeasure(id))
}
