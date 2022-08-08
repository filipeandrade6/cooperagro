package unitofmeasure

import (
	"strings"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type inmem struct {
	m map[entity.ID]*entity.UnitOfMeasure
}

func newInmem() *inmem {
	return &inmem{
		m: map[entity.ID]*entity.UnitOfMeasure{},
	}
}

func (i *inmem) GetUnitOfMeasureByID(id entity.ID) (*entity.UnitOfMeasure, error) {
	bp, ok := i.m[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	return bp, nil
}

func (i *inmem) SearchUnitOfMeasure(query string) ([]*entity.UnitOfMeasure, error) {
	var d []*entity.UnitOfMeasure
	for _, j := range i.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}

	return d, nil
}

func (i *inmem) ListUnitOfMeasure() ([]*entity.UnitOfMeasure, error) {
	var d []*entity.UnitOfMeasure
	for _, j := range i.m {
		d = append(d, j)
	}

	return d, nil
}

func (i *inmem) CreateUnitOfMeasure(e *entity.UnitOfMeasure) (entity.ID, error) {
	i.m[e.ID] = e

	return e.ID, nil
}

func (i *inmem) UpdateUnitOfMeasure(e *entity.UnitOfMeasure) error {
	_, err := i.GetUnitOfMeasureByID(e.ID)
	if err != nil {
		return err
	}

	i.m[e.ID] = e

	return nil
}

func (i *inmem) DeleteUnitOfMeasure(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}

	delete(i.m, id)

	return nil
}
