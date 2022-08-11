package inventory

import "github.com/filipeandrade6/cooperagro/domain/entity"

// TODO porquer duas interfaces e pq interface Reader e Writer separada?

// Driver Adapter (preciso que o repositorio consiga fazer isso)

type Reader interface {
	GetInventoryByID(id entity.ID) (*entity.Inventory, error)
	ListInventory() ([]*entity.Inventory, error)
}

type Writer interface {
	CreateInventory(e *entity.Inventory) (entity.ID, error)
	UpdateInventory(e *entity.Inventory) error
	DeleteInventory(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

// Driven Adapter (o que o use case comanda)

type UseCase interface {
	GetInventoryByID(id entity.ID) (*entity.Inventory, error)
	ListInventory() ([]*entity.Inventory, error)
	CreateInventory(
		customerID,
		productID entity.ID,
		quantity int,
		unitOfMeasureID entity.ID,
	) (entity.ID, error)
	UpdateInventory(e *entity.Inventory) error
	DeleteInventory(id entity.ID) error
}
