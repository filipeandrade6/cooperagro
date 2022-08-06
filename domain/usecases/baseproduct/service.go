package baseproduct

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

func (s *Service) GetByID(id entities.ID) (*entities.BaseProduct, error) {
	bp, err := s.repo.GetBaseProductByID(id)
	if bp == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return bp, nil
}

func (s *Service) Search(query string) ([]*entities.BaseProduct, error) {
	baseProducts, err := s.repo.SearchBaseProduct(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(baseProducts) == 0 {
		return nil, entities.ErrNotFound
	}

	return baseProducts, nil
}

func (s *Service) List() ([]*entities.BaseProduct, error) {
	baseProducts, err := s.repo.ListBaseProduct()
	if err != nil {
		return nil, err
	}
	if len(baseProducts) == 0 {
		return nil, entities.ErrNotFound
	}

	return baseProducts, nil
}

func (s *Service) Create(name string) (entities.ID, error) {
	bp, err := entities.NewBaseProduct(name)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.CreateBaseProduct(bp)
}

func (s *Service) Update(e *entities.BaseProduct) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateBaseProduct(e)
}

func (s *Service) Delete(id entities.ID) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}

	return s.repo.DeleteBaseProduct(id)
}
