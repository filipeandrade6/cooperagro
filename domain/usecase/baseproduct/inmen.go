package baseproduct

import (
	"strings"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type inmem struct {
	m map[entity.ID]*entity.BaseProduct
}

func newInmem() *inmem {
	return &inmem{
		m: map[entity.ID]*entity.BaseProduct{},
	}
}

func (i *inmem) GetBaseProductByID(id entity.ID) (*entity.BaseProduct, error) {
	bp, ok := i.m[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	return bp, nil
}

func (i *inmem) SearchBaseProduct(query string) ([]*entity.BaseProduct, error) {
	var d []*entity.BaseProduct
	for _, j := range i.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}

	return d, nil
}

func (i *inmem) ListBaseProduct() ([]*entity.BaseProduct, error) {
	var d []*entity.BaseProduct
	for _, j := range i.m {
		d = append(d, j)
	}

	return d, nil
}

func (i *inmem) CreateBaseProduct(e *entity.BaseProduct) (entity.ID, error) {
	for _, j := range i.m {
		if e.Name == j.Name {
			return e.ID, entity.ErrEntityAlreadyExists
		}
	}

	i.m[e.ID] = e

	return e.ID, nil
}

func (i *inmem) UpdateBaseProduct(e *entity.BaseProduct) error {
	_, err := i.GetBaseProductByID(e.ID)
	if err != nil {
		return err
	}

	for _, j := range i.m {
		if e.Name == j.Name && e.ID != j.ID {
			return entity.ErrEntityAlreadyExists
		}
	}

	i.m[e.ID] = e

	return nil
}

func (i *inmem) DeleteBaseProduct(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}

	delete(i.m, id)

	return nil
}
