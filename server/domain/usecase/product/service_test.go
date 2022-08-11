package product

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/product/mock"
	"github.com/stretchr/testify/assert"
)

func newFixtureProduct(validBaseProductID entity.ID) *entity.Product {
	return &entity.Product{
		Name:          "lima",
		BaseProductID: validBaseProductID,
	}
}

func TestService_GetProductByID(t *testing.T) {
	bpID := entity.NewID()
	bp := mock.NewMockBaseProductService(bpID)
	repo := newInmem()
	s := NewService(bp, repo)
	p := newFixtureProduct(bpID)
	id, err := s.CreateProduct(p.Name, p.BaseProductID)
	assert.Nil(t, err)

	t.Run("get existent", func(t *testing.T) {
		saved, err := s.GetProductByID(id)
		assert.Nil(t, err)
		assert.Equal(t, p.Name, saved.Name)
	})

	t.Run("get non existent", func(t *testing.T) {
		_, err := s.GetProductByID(entity.NewID())
		assert.Equal(t, entity.ErrNotFound, err)
	})
}

func TestService_SearchProduct(t *testing.T) {
	bpID := entity.NewID()
	bp := mock.NewMockBaseProductService(bpID)
	repo := newInmem()
	s := NewService(bp, repo)
	p := newFixtureProduct(bpID)

	_, _ = s.CreateProduct(p.Name, p.BaseProductID)

	t.Run("search equal", func(t *testing.T) {
		saved, err := s.SearchProduct("lima")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, p.Name, saved[0].Name)
	})

	t.Run("search equal but with capital letters", func(t *testing.T) {
		saved, err := s.SearchProduct("LimA")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(saved))
		assert.Equal(t, p.Name, saved[0].Name)
	})

	t.Run("search for inexistent", func(t *testing.T) {
		saved, err := s.SearchProduct("baía")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(saved))
	})
}

func TestService_ListProduct(t *testing.T) {
	bpID := entity.NewID()
	bp := mock.NewMockBaseProductService(bpID)
	repo := newInmem()
	s := NewService(bp, repo)

	t.Run("list empty", func(t *testing.T) {
		ps, err := s.ListProduct()
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, 0, len(ps))
	})

	t.Run("list all", func(t *testing.T) {
		p1 := newFixtureProduct(bpID)
		p2 := newFixtureProduct(bpID)
		p2.Name = "baía"
		_, _ = s.CreateProduct(p1.Name, p1.BaseProductID)
		_, _ = s.CreateProduct(p2.Name, p2.BaseProductID)

		ps, err := s.ListProduct()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(ps))
	})
}

func TestService_CreateProduct(t *testing.T) {
	bpID := entity.NewID()
	bp := mock.NewMockBaseProductService(bpID)
	repo := newInmem()
	s := NewService(bp, repo)
	p := newFixtureProduct(bpID)

	t.Run("create product with invalid base product", func(t *testing.T) {
		_, err := s.CreateProduct(p.Name, entity.NewID())
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("create product", func(t *testing.T) {
		_, err := s.CreateProduct(p.Name, p.BaseProductID)
		assert.Nil(t, err)
	})

	t.Run("create existent product", func(t *testing.T) {
		_, err := s.CreateProduct(p.Name, p.BaseProductID)
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_UpdateProduct(t *testing.T) {
	bpID := entity.NewID()
	bp := mock.NewMockBaseProductService(bpID)
	repo := newInmem()
	s := NewService(bp, repo)

	p := newFixtureProduct(bpID)
	p2 := newFixtureProduct(bpID)
	p2.Name = "baía"

	id, _ := s.CreateProduct(p.Name, p.BaseProductID)
	_, _ = s.CreateProduct(p2.Name, p.BaseProductID)

	t.Run("update product with invalid base product", func(t *testing.T) {
		err := s.UpdateProduct(&entity.Product{
			ID:            id,
			Name:          "natal",
			BaseProductID: entity.NewID(),
		})
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("update product", func(t *testing.T) {
		err := s.UpdateProduct(&entity.Product{
			ID:            id,
			Name:          "natal",
			BaseProductID: p.BaseProductID,
		})
		assert.Nil(t, err)
	})

	t.Run("update to existent product", func(t *testing.T) {
		err := s.UpdateProduct(&entity.Product{
			ID:            id,
			Name:          p2.Name,
			BaseProductID: p.BaseProductID,
		})
		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
	})
}

func TestService_DeleteProduct(t *testing.T) {
	bpID := entity.NewID()
	bp := mock.NewMockBaseProductService(bpID)
	repo := newInmem()
	s := NewService(bp, repo)

	p := newFixtureProduct(bpID)
	id, _ := s.CreateProduct(p.Name, p.BaseProductID)

	t.Run("delete product", func(t *testing.T) {
		err := s.DeleteProduct(id)
		assert.Nil(t, err)
	})

	t.Run("delete inexistent product", func(t *testing.T) {
		err := s.DeleteProduct(id)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}
