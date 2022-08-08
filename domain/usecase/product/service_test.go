package product

import (
	"testing"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureProduct() *entity.Product {
	return &entity.Product{
		Name:          "lima",
		BaseProductID: entity.NewID(),
		CreatedAt:     time.Now(),
	}
}

func TestCreate(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureProduct()
	_, err := s.CreateProduct(bp.Name, bp.BaseProductID)
	assert.Nil(t, err)
	assert.False(t, bp.CreatedAt.IsZero())
}

func TestSearchListGetProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp1 := newFixtureProduct()
	bp2 := newFixtureProduct()
	bp2.Name = "bahia"

	uID, _ := s.CreateProduct(bp1.Name, bp1.BaseProductID)
	_, _ = s.CreateProduct(bp2.Name, bp2.BaseProductID)

	t.Run("search", func(t *testing.T) {
		bp, err := s.SearchProduct("BAHIA")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(bp))
		assert.Equal(t, "bahia", bp[0].Name)

		bp, err = s.SearchProduct("pera")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, bp)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := s.ListProduct()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		bp, err := s.GetProductByID(uID)
		assert.Nil(t, err)
		assert.Equal(t, bp.Name, bp.Name)
	})
}

// TODO teste update para nome existente (que deveria ser unico)

func TestUpdateDeleteProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	bp := newFixtureProduct()

	id, err := s.CreateProduct(bp.Name, bp.BaseProductID)
	assert.Nil(t, err)
	saved, _ := s.GetProductByID(id)
	saved.Name = "bahia"
	assert.Nil(t, s.UpdateProduct(saved))
	updated, err := s.GetProductByID(id)
	assert.Nil(t, err)
	assert.Equal(t, "bahia", updated.Name)

	assert.Nil(t, s.DeleteProduct(id))
	assert.Equal(t, entity.ErrNotFound, s.DeleteProduct(id))
}
