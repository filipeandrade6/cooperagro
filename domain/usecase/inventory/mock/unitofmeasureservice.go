package mock

import (
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type MockUnitOfMeasureService struct {
	validID entity.ID
}

func NewMockUnitOfMeasureService(validID entity.ID) MockUnitOfMeasureService {
	return MockUnitOfMeasureService{validID: validID}
}

func (m MockUnitOfMeasureService) GetUnitOfMeasureByID(id entity.ID) (*entity.UnitOfMeasure, error) {
	if id == m.validID {
		return &entity.UnitOfMeasure{
			ID:        id,
			Name:      "kilogram",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	return nil, entity.ErrNotFound
}

func (m MockUnitOfMeasureService) SearchUnitOfMeasure(query string) ([]*entity.UnitOfMeasure, error) {
	return nil, nil
}

func (m MockUnitOfMeasureService) ListUnitOfMeasure() ([]*entity.UnitOfMeasure, error) {
	return nil, nil
}

func (m MockUnitOfMeasureService) CreateUnitOfMeasure(name string) (entity.ID, error) {
	return entity.NewID(), nil
}

func (m MockUnitOfMeasureService) UpdateUnitOfMeasure(e *entity.UnitOfMeasure) error {
	return nil
}

func (m MockUnitOfMeasureService) DeleteUnitOfMeasure(id entity.ID) error {
	return nil
}
