package product

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

func (s *Service) GetByID(id entities.ID) (*entities.Product, error) {
	p, err := s.repo.GetByID(id)
	if p == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Service) Search(query string) ([]*entities.Product, error) {
	products, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entities.ErrNotFound
	}

	return products, nil
}

func (s *Service) List() ([]*entities.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entities.ErrNotFound
	}

	return products, nil
}

func (s *Service) Create(name string, baseProduct entities.ID) (entities.ID, error) {
	p, err := entities.NewProduct(name, baseProduct)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.Create(p)
}

func (s *Service) Update(e *entities.Product) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.Update(e)
}

func (s *Service) Delete(id entities.ID) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}

	return s.repo.Delete(id)
}
