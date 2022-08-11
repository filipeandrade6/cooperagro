package baseproduct

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

func (s *Service) GetBaseProductByID(id entity.ID) (*entity.BaseProduct, error) {
	bp, err := s.repo.GetBaseProductByID(id)
	if bp == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return bp, nil
}

func (s *Service) SearchBaseProduct(query string) ([]*entity.BaseProduct, error) {
	baseProducts, err := s.repo.SearchBaseProduct(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(baseProducts) == 0 {
		return nil, entity.ErrNotFound
	}

	return baseProducts, nil
}

func (s *Service) ListBaseProduct() ([]*entity.BaseProduct, error) {
	baseProducts, err := s.repo.ListBaseProduct()
	if err != nil {
		return nil, err
	}
	if len(baseProducts) == 0 {
		return nil, entity.ErrNotFound
	}

	return baseProducts, nil
}

func (s *Service) CreateBaseProduct(name string) (entity.ID, error) {
	bp, err := entity.NewBaseProduct(name)
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.CreateBaseProduct(bp)
}

func (s *Service) UpdateBaseProduct(e *entity.BaseProduct) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateBaseProduct(e)
}

func (s *Service) DeleteBaseProduct(id entity.ID) error {
	if _, err := s.GetBaseProductByID(id); err != nil {
		return err
	}

	return s.repo.DeleteBaseProduct(id)
}
