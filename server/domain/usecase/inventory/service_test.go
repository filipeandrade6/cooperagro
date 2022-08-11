package inventory

// import (
// 	"testing"

// 	"github.com/filipeandrade6/cooperagro/domain/entity"
// 	"github.com/filipeandrade6/cooperagro/domain/usecase/inventory/mock"
// 	"github.com/stretchr/testify/assert"
// )

// func newFixtureInventory(validProductID, validUnitOfMeasureID, validUserID entity.ID) *entity.Inventory {
// 	return &entity.Inventory{
// 		UserID:          validUserID,
// 		ProductID:       validProductID,
// 		Quantity:        23,
// 		UnitOfMeasureID: validUnitOfMeasureID,
// 	}
// }

// func TestService_GetInventoryByID(t *testing.T) {
// 	repo := newInmem()

// 	pID := entity.NewID()
// 	umID := entity.NewID()
// 	uID := entity.NewID()

// 	p := mock.NewMockProductService(pID)
// 	um := mock.NewMockUnitOfMeasureService(umID)
// 	u := mock.NewMockUserService(uID)

// 	s := NewService(p, um, u, repo)

// 	i := newFixtureInventory(pID, umID, uID)
// 	id, err := s.CreateInventory(
// 		i.UserID,
// 		i.ProductID,
// 		i.Quantity,
// 		i.UnitOfMeasureID,
// 	)
// 	assert.Nil(t, err)

// 	t.Run("get existent", func(t *testing.T) {
// 		saved, err := s.GetInventoryByID(id)
// 		assert.Nil(t, err)
// 		assert.Equal(t, i.UserID, saved.UserID)
// 		assert.Equal(t, i.ProductID, saved.ProductID)
// 		assert.Equal(t, i.Quantity, saved.Quantity)
// 		assert.Equal(t, i.UnitOfMeasureID, saved.UnitOfMeasureID)
// 	})

// 	t.Run("get non existent", func(t *testing.T) {
// 		_, err := s.GetInventoryByID(entity.NewID())
// 		assert.Equal(t, entity.ErrNotFound, err)
// 	})
// }

// func TestService_ListInventory(t *testing.T) {
// 	repo := newInmem()

// 	pID := entity.NewID()
// 	umID := entity.NewID()
// 	uID := entity.NewID()

// 	p := mock.NewMockProductService(pID)
// 	um := mock.NewMockUnitOfMeasureService(umID)
// 	u := mock.NewMockUserService(uID)

// 	s := NewService(p, um, u, repo)

// 	t.Run("list empty", func(t *testing.T) {
// 		is, err := s.ListInventory()
// 		assert.Equal(t, entity.ErrNotFound, err)
// 		assert.Equal(t, 0, len(is))
// 	})

// 	t.Run("list all", func(t *testing.T) {
// 		i1 := newFixtureInventory(pID, umID, uID)
// 		i2 := newFixtureInventory(pID, umID, uID)
// 		i2.ProductID = entity.NewID()
// 		_, _ = s.CreateInventory(
// 			i1.UserID,
// 			i1.ProductID,
// 			i1.Quantity,
// 			i1.UnitOfMeasureID,
// 		)
// 		_, _ = s.CreateInventory(
// 			i2.UserID,
// 			i2.ProductID,
// 			i2.Quantity,
// 			i2.UnitOfMeasureID,
// 		)

// 		is, err := s.ListInventory()
// 		assert.Nil(t, err)
// 		assert.Equal(t, 2, len(is))
// 	})
// }

// func TestService_CreateInventory(t *testing.T) {
// 	repo := newInmem()
// 	s := NewService(repo)
// 	i := newFixtureInventory()

// 	t.Run("create inventory", func(t *testing.T) {
// 		_, err := s.CreateInventory(
// 			i.UserID,
// 			i.ProductID,
// 			i.Quantity,
// 			i.UnitOfMeasureID,
// 		)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("create existent inventory", func(t *testing.T) {
// 		_, err := s.CreateInventory(
// 			i.UserID,
// 			i.ProductID,
// 			111,
// 			i.UnitOfMeasureID,
// 		)
// 		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
// 	})
// }

// func TestService_UpdateInventory(t *testing.T) {
// 	repo := newInmem()
// 	s := NewService(repo)

// 	i1 := newFixtureInventory()
// 	i2 := newFixtureInventory()
// 	i2.UnitOfMeasureID = entity.NewID()

// 	id, _ := s.CreateInventory(
// 		i1.UserID,
// 		i1.ProductID,
// 		i1.Quantity,
// 		i1.UnitOfMeasureID,
// 	)
// 	_, _ = s.CreateInventory(
// 		i2.UserID,
// 		i2.ProductID,
// 		i2.Quantity,
// 		i2.UnitOfMeasureID,
// 	)

// 	t.Run("update inventory", func(t *testing.T) {
// 		err := s.UpdateInventory(&entity.Inventory{
// 			ID:              id,
// 			UserID:          i1.UserID,
// 			ProductID:       entity.NewID(),
// 			Quantity:        25,
// 			UnitOfMeasureID: i1.UnitOfMeasureID,
// 		})
// 		assert.Nil(t, err)
// 	})

// 	t.Run("update to existent inventory", func(t *testing.T) {
// 		err := s.UpdateInventory(&entity.Inventory{
// 			ID:              id,
// 			UserID:          i1.UserID,
// 			ProductID:       i1.ProductID,
// 			Quantity:        100,
// 			UnitOfMeasureID: i2.UnitOfMeasureID,
// 		})
// 		assert.Equal(t, entity.ErrEntityAlreadyExists, err)
// 	})
// }

// func TestService_DeleteInventory(t *testing.T) {
// 	repo := newInmem()
// 	s := NewService(repo)
// 	i := newFixtureInventory()
// 	id, _ := s.CreateInventory(
// 		i.UserID,
// 		i.ProductID,
// 		i.Quantity,
// 		i.UnitOfMeasureID,
// 	)

// 	t.Run("delete base product", func(t *testing.T) {
// 		err := s.DeleteInventory(id)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("delete inexistent base product", func(t *testing.T) {
// 		err := s.DeleteInventory(id)
// 		assert.Equal(t, entity.ErrNotFound, err)
// 	})
// }
