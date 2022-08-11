package inventory

import (
	"errors"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/filipeandrade6/cooperagro/domain/usecase/product"
	"github.com/filipeandrade6/cooperagro/domain/usecase/unitofmeasure"
	"github.com/filipeandrade6/cooperagro/domain/usecase/user"
)

type Service struct {
	productService       product.UseCase
	unitOfMeasureService unitofmeasure.UseCase
	userService          user.UseCase
	repo                 Repository
}

func NewService(p product.UseCase, um unitofmeasure.UseCase, u user.UseCase, r Repository) *Service {
	return &Service{
		productService:       p,
		unitOfMeasureService: um,
		userService:          u,
		repo:                 r,
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
	userID,
	productID entity.ID,
	quantity int,
	unitOfMeasureID entity.ID,
) (entity.ID, error) {
	_, err := s.productService.GetProductByID(productID)
	if errors.Is(err, entity.ErrNotFound) {
		return entity.NewID(), err
	}

	_, err = s.userService.GetUserByID(userID)
	if errors.Is(err, entity.ErrNotFound) {
		return entity.NewID(), err
	}

	i, err := entity.NewInventory(
		userID,
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
	_, err := s.productService.GetProductByID(e.ProductID)
	if errors.Is(err, entity.ErrNotFound) {
		return err
	}

	_, err = s.userService.GetUserByID(e.UserID)
	if errors.Is(err, entity.ErrNotFound) {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateInventory(e)
}

func (s *Service) DeleteInventory(id entity.ID) error {
	if _, err := s.GetInventoryByID(id); err != nil {
		return err
	}

	return s.repo.DeleteInventory(id)
}
