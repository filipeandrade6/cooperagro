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

func (s *Service) GetProductByID(id entities.ID) (*entities.Product, error) {
	p, err := s.repo.GetProductByID(id)
	if p == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Service) SearchProduct(query string) ([]*entities.Product, error) {
	products, err := s.repo.SearchProduct(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entities.ErrNotFound
	}

	return products, nil
}

func (s *Service) ListProduct() ([]*entities.Product, error) {
	products, err := s.repo.ListProduct()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entities.ErrNotFound
	}

	return products, nil
}

func (s *Service) CreateProduct(name string, baseProductID entities.ID) (entities.ID, error) {
	p, err := entities.NewProduct(name, baseProductID)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.CreateProduct(p)
}

func (s *Service) UpdateProduct(e *entities.Product) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateProduct(e)
}

func (s *Service) DeleteProduct(id entities.ID) error {
	if _, err := s.GetProductByID(id); err != nil {
		return err
	}

	return s.repo.DeleteProduct(id)
}
