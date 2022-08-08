package inventory

import (
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

func (s *Service) GetInventoryByID(id entity.ID) (*entity.Inventory, error) {
	i, err := s.repo.GetInventoryByID(id)
	if i == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (s *Service) ListInventory() ([]*entity.Inventory, error) {
	inventories, err := s.repo.ListInventory()
	if err != nil {
		return nil, err
	}
	if len(inventories) == 0 {
		return nil, entity.ErrNotFound
	}

	return inventories, nil
}

func (s *Service) CreateInventory(
	customerID,
	productID entity.ID,
	quantity int,
	unitOfMeasureID entity.ID,
) (entity.ID, error) {
	i, err := entity.NewInventory(
		customerID,
		productID,
		quantity,
		unitOfMeasureID,
	)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.CreateInventory(i)
}

func (s *Service) UpdateInventory(e *entity.Inventory) error {
	e.UpdatedAt = time.Now()

	return s.repo.UpdateInventory(e)
}

func (s *Service) DeleteInventory(id entity.ID) error {
	if _, err := s.GetInventoryByID(id); err != nil {
		return err
	}

	return s.repo.DeleteInventory(id)
}
