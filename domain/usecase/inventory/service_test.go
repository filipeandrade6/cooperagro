package inventory

import (
	"testing"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureInventory() *entity.Inventory {
	return &entity.Inventory{
		ID:              entity.NewID(),
		UserID:          entity.NewID(),
		ProductID:       entity.NewID(),
		Quantity:        23,
		UnitOfMeasureID: entity.NewID(),
		CreatedAt:       time.Now(),
	}
}

func TestCreate(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	i := newFixtureInventory()
	_, err := s.CreateInventory(i.UserID, i.ProductID, i.Quantity, i.UnitOfMeasureID)
	assert.Nil(t, err)
	assert.False(t, i.CreatedAt.IsZero())
}

func TestListGetInventory(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	i := newFixtureInventory()

	uID, _ := s.CreateInventory(i.UserID, i.ProductID, i.Quantity, i.UnitOfMeasureID)

	t.Run("list all", func(t *testing.T) {
		all, err := s.ListInventory()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(all))
	})

	t.Run("get", func(t *testing.T) {
		j, err := s.GetInventoryByID(uID)
		assert.Nil(t, err)
		assert.Equal(t, i.Quantity, j.Quantity)
	})
}

func TestUpdateDeleteInventory(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	i := newFixtureInventory()

	id, err := s.CreateInventory(i.UserID, i.ProductID, i.Quantity, i.UnitOfMeasureID)
	assert.Nil(t, err)
	saved, _ := s.GetInventoryByID(id)
	saved.Quantity = 42
	assert.Nil(t, s.UpdateInventory(saved))
	updated, err := s.GetInventoryByID(id)
	assert.Nil(t, err)
	assert.Equal(t, 42, updated.Quantity)

	assert.Nil(t, s.DeleteInventory(id))
	assert.Equal(t, entity.ErrNotFound, s.DeleteInventory(id))
}
