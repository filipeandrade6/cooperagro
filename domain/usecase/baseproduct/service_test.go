package baseproduct

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureBaseProduct() *entity.BaseProduct {
	return &entity.BaseProduct{
		Name: "tomate",
	}
}

func TestService_GetBaseProductByID(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureBaseProduct()
	id, err := s.CreateBaseProduct(bp.Name)
	assert.Nil(t, err)

	t.Run("get existent", func(t *testing.T) {
		saved, err := s.GetBaseProductByID(id)
		assert.Nil(t, err)
		assert.Equal(t, bp.Name, saved.Name)
	})

	t.Run("get non existent", func(t *testing.T) {
		_, err := s.GetBaseProductByID(entity.NewID())
		assert.Equal(t, entity.ErrNotFound, err)
	})
}

func TestService_SearchBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureBaseProduct()

	_, _ = s.CreateBaseProduct(bp.Name)

	t.Run("search equal", func(t *testing.T) {
		saved, err := s.SearchBaseProduct("tomate")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, bp.Name, saved[0].Name)
	})

	t.Run("search equal but with capital letters", func(t *testing.T) {
		saved, err := s.SearchBaseProduct("TomatE")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, bp.Name, saved[0].Name)
	})

	t.Run("search for inexistent", func(t *testing.T) {
		saved, err := s.SearchBaseProduct("morango")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(saved))
	})
}

func TestService_ListBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)

	t.Run("list empty", func(t *testing.T) {
		bps, err := s.ListBaseProduct()
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(bps))
	})

	t.Run("list all", func(t *testing.T) {
		bp1 := newFixtureBaseProduct()
		bp2 := newFixtureBaseProduct()
		bp2.Name = "acerola"
		_, _ = s.CreateBaseProduct(bp1.Name)
		_, _ = s.CreateBaseProduct(bp2.Name)

		bps, err := s.ListBaseProduct()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(bps))
	})
}

func TestService_CreateBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureBaseProduct()

	t.Run("create base product", func(t *testing.T) {
		_, err := s.CreateBaseProduct(bp.Name)
		assert.Nil(t, err)
	})

	t.Run("create existent base product", func(t *testing.T) {
		_, err := s.CreateBaseProduct(bp.Name)
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_UpdateBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)

	bp := newFixtureBaseProduct()
	bp2 := newFixtureBaseProduct()
	bp2.Name = "manga"

	id, _ := s.CreateBaseProduct(bp.Name)
	_, _ = s.CreateBaseProduct(bp2.Name)

	t.Run("update base product", func(t *testing.T) {
		err := s.UpdateBaseProduct(&entity.BaseProduct{
			ID:   id,
			Name: "acerola",
		})
		assert.Nil(t, err)
	})

	t.Run("update to existent base product", func(t *testing.T) {
		err := s.UpdateBaseProduct(&entity.BaseProduct{
			ID:   id,
			Name: bp2.Name,
		})
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_DeleteBaseProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureBaseProduct()
	id, _ := s.CreateBaseProduct(bp.Name)

	t.Run("delete base product", func(t *testing.T) {
		err := s.DeleteBaseProduct(id)
		assert.Nil(t, err)
	})

	t.Run("delete inexistent base product", func(t *testing.T) {
		err := s.DeleteBaseProduct(id)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}
