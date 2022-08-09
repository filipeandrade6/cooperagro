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
	p1 := newFixtureProduct()
	p2 := newFixtureProduct()
	p2.Name = "bahia"

	uID, _ := s.CreateProduct(p1.Name, p1.BaseProductID)
	_, _ = s.CreateProduct(p2.Name, p2.BaseProductID)

	t.Run("search", func(t *testing.T) {
		p, err := s.SearchProduct("BAHIA")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(p))
		assert.Equal(t, "bahia", p[0].Name)

		p, err = s.SearchProduct("pera")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, p)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := s.ListProduct()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		p, err := s.GetProductByID(uID)
		assert.Nil(t, err)
		assert.Equal(t, p.Name, p.Name)
	})
}

// TODO teste update para nome existente (que deveria ser unico)

func TestUpdateDeleteProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	p := newFixtureProduct()

	id, err := s.CreateProduct(p.Name, p.BaseProductID)
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

func TestCreateUpdateSameNameAndBaseProductIDProduct(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	p := newFixtureProduct()

	_, err := s.CreateProduct(p.Name, p.BaseProductID)
	assert.Nil(t, err)
	_, err = s.CreateProduct(p.Name, p.BaseProductID)
	assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	assert.Equal(t, entity.ErrEntityAlreadyExists, s.UpdateProduct(p))
}
