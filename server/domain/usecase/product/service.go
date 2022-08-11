package product

import (
	"errors"
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/baseproduct"
)

type Service struct {
	baseProductService baseproduct.UseCase
	repo               Repository
}

func NewService(b baseproduct.UseCase, r Repository) *Service {
	return &Service{
		baseProductService: b,
		repo:               r,
	}
}

func (s *Service) GetProductByID(id entity.ID) (*entity.Product, error) {
	p, err := s.repo.GetProductByID(id)
	if p == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Service) SearchProduct(query string) ([]*entity.Product, error) {
	products, err := s.repo.SearchProduct(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}

	return products, nil
}

func (s *Service) ListProduct() ([]*entity.Product, error) {
	products, err := s.repo.ListProduct()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}

	return products, nil
}

func (s *Service) CreateProduct(name string, baseProductID entity.ID) (entity.ID, error) {
	_, err := s.baseProductService.GetBaseProductByID(baseProductID)
	if errors.Is(err, entity.ErrNotFound) {
		return entity.NewID(), err
	}

	p, err := entity.NewProduct(name, baseProductID)
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.CreateProduct(p)
}

func (s *Service) UpdateProduct(e *entity.Product) error {
	_, err := s.baseProductService.GetBaseProductByID(e.BaseProductID)
	if errors.Is(err, entity.ErrNotFound) {
		return err
	}

	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateProduct(e)
}

func (s *Service) DeleteProduct(id entity.ID) error {
	if _, err := s.GetProductByID(id); err != nil {
		return err
	}

	return s.repo.DeleteProduct(id)
}
