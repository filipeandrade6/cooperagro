package unitofmeasure

import (
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entities"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUnitOfMeasureByID(id entities.ID) (*entities.UnitOfMeasure, error) {
	u, err := s.repo.GetUnitOfMeasureByID(id)
	if u == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) SearchUnitOfMeasure(query string) ([]*entities.UnitOfMeasure, error) {
	unitsOfMeasure, err := s.repo.SearchUnitOfMeasure(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(unitsOfMeasure) == 0 {
		return nil, entities.ErrNotFound
	}

	return unitsOfMeasure, nil
}

func (s *Service) ListUnitOfMeasure() ([]*entities.UnitOfMeasure, error) {
	unitsOfMeasure, err := s.repo.ListUnitOfMeasure()
	if err != nil {
		return nil, err
	}
	if len(unitsOfMeasure) == 0 {
		return nil, entities.ErrNotFound
	}

	return unitsOfMeasure, nil
}

func (s *Service) CreateUnitOfMeasure(name string) (entities.ID, error) {
	u, err := entities.NewUnitOfMeasure(name)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.CreateUnitOfMeasure(u)
}

func (s *Service) UpdateUnitOfMeasure(e *entities.UnitOfMeasure) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateUnitOfMeasure(e)
}

func (s *Service) DeleteUnitOfMeasure(id entities.ID) error {
	if _, err := s.GetUnitOfMeasureByID(id); err != nil {
		return err
	}

	return s.repo.DeleteUnitOfMeasure(id)
}
