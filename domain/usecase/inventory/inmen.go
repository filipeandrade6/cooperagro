package inventory

import (
	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type inmem struct {
	m map[entity.ID]*entity.Inventory
}

func newInmem() *inmem {
	return &inmem{
		m: map[entity.ID]*entity.Inventory{},
	}
}

func (i *inmem) GetInventoryByID(id entity.ID) (*entity.Inventory, error) {
	bp, ok := i.m[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	return bp, nil
}

func (i *inmem) ListInventory() ([]*entity.Inventory, error) {
	var d []*entity.Inventory
	for _, j := range i.m {
		d = append(d, j)
	}

	return d, nil
}

func (i *inmem) CreateInventory(e *entity.Inventory) (entity.ID, error) {
	for _, j := range i.m {
		if e.UserID == j.UserID && e.ProductID == j.ProductID && e.UnitOfMeasureID == j.UnitOfMeasureID {
			return e.ID, entity.ErrEntityAlreadyExists
		}
	}

	i.m[e.ID] = e

	return e.ID, nil
}

func (i *inmem) UpdateInventory(e *entity.Inventory) error {
	_, err := i.GetInventoryByID(e.ID)
	if err != nil {
		return err
	}

	for _, j := range i.m {
		if e.UserID == j.UserID && e.ProductID == j.ProductID && e.UnitOfMeasureID == j.UnitOfMeasureID {
			return entity.ErrEntityAlreadyExists
		}
	}

	i.m[e.ID] = e

	return nil
}

func (i *inmem) DeleteInventory(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}

	delete(i.m, id)

	return nil
}
