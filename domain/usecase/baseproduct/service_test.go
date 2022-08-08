package baseproduct

import (
	"testing"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureBaseProduct() *entity.BaseProduct {
	return &entity.BaseProduct{
		Name:      "tomate",
		CreatedAt: time.Now(),
	}
}

func TestCreate(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureBaseProduct()
	_, err := s.CreateBaseProduct((bp.Name))
	assert.Nil(t, err)
	assert.False(t, bp.CreatedAt.IsZero())
}

func TestSearchListGetBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp1 := newFixtureBaseProduct()
	bp2 := newFixtureBaseProduct()
	bp2.Name = "manga"

	uID, _ := s.CreateBaseProduct(bp1.Name)
	_, _ = s.CreateBaseProduct(bp2.Name)

	t.Run("search", func(t *testing.T) {
		bp, err := s.SearchBaseProduct("MANGA")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(bp))
		assert.Equal(t, "manga", bp[0].Name)

		bp, err = s.SearchBaseProduct("acerola")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, bp)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := s.ListBaseProduct()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		bp, err := s.GetBaseProductByID(uID)
		assert.Nil(t, err)
		assert.Equal(t, bp1.Name, bp.Name)
	})
}

// TODO teste update para nome existente (que deveria ser unico)

func TestUpdateDeleteBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureBaseProduct()

	id, err := s.CreateBaseProduct(bp.Name)
	assert.Nil(t, err)
	saved, _ := s.GetBaseProductByID(id)
	saved.Name = "manga"
	assert.Nil(t, s.UpdateBaseProduct(saved))
	updated, err := s.GetBaseProductByID(id)
	assert.Nil(t, err)
	assert.Equal(t, "manga", updated.Name)

	assert.Nil(t, s.DeleteBaseProduct(id))
	assert.Equal(t, entity.ErrNotFound, s.DeleteBaseProduct(id))
}
