package inventory

import (
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

func (s *Service) GetInventoryByID(id entities.ID) (*entities.Inventory, error) {
	i, err := s.repo.GetInventoryByID(id)
	if i == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (s *Service) ListInventory() ([]*entities.Inventory, error) {
	inventories, err := s.repo.ListInventory()
	if err != nil {
		return nil, err
	}
	if len(inventories) == 0 {
		return nil, entities.ErrNotFound
	}

	return inventories, nil
}

func (s *Service) CreateInventory(
	customerID,
	productID entities.ID,
	quantity int,
	unitOfMeasureID entities.ID,
) (entities.ID, error) {
	i := entities.NewInventory(
		customerID,
		productID,
		quantity,
		unitOfMeasureID,
	)
	return s.repo.CreateInventory(i)
}

func (s *Service) UpdateInventory(e *entities.Inventory) error {
	e.UpdatedAt = time.Now()

	return s.repo.UpdateInventory(e)
}

func (s *Service) DeleteInventory(id entities.ID) error {
	if _, err := s.GetInventoryByID(id); err != nil {
		return err
	}

	return s.repo.DeleteInventory(id)
}
