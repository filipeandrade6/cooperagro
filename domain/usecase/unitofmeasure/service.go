package unitofmeasure

import (
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUnitOfMeasureByID(id entity.ID) (*entity.UnitOfMeasure, error) {
	u, err := s.repo.GetUnitOfMeasureByID(id)
	if u == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) SearchUnitOfMeasure(query string) ([]*entity.UnitOfMeasure, error) {
	unitsOfMeasure, err := s.repo.SearchUnitOfMeasure(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(unitsOfMeasure) == 0 {
		return nil, entity.ErrNotFound
	}

	return unitsOfMeasure, nil
}

func (s *Service) ListUnitOfMeasure() ([]*entity.UnitOfMeasure, error) {
	unitsOfMeasure, err := s.repo.ListUnitOfMeasure()
	if err != nil {
		return nil, err
	}
	if len(unitsOfMeasure) == 0 {
		return nil, entity.ErrNotFound
	}

	return unitsOfMeasure, nil
}

func (s *Service) CreateUnitOfMeasure(name string) (entity.ID, error) {
	u, err := entity.NewUnitOfMeasure(name)
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.CreateUnitOfMeasure(u)
}

func (s *Service) UpdateUnitOfMeasure(e *entity.UnitOfMeasure) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateUnitOfMeasure(e)
}

func (s *Service) DeleteUnitOfMeasure(id entity.ID) error {
	if _, err := s.GetUnitOfMeasureByID(id); err != nil {
		return err
	}

	return s.repo.DeleteUnitOfMeasure(id)
}
