package inventory

import "github.com/filipeandrade6/cooperagro/src/domain/entities"

// TODO porquer duas interfaces e pq interface Reader e Writer separada?

type Reader interface {
	GetByID(id entities.ID) (*entities.Inventory, error)
	Search(query string) ([]*entities.Inventory, error)
	List() ([]*entities.Inventory, error)
}

type Writer interface {
	Create(e *entities.Inventory) entities.ID
	Update(e *entities.Inventory) error
	Delete(id entities.ID) error
}

type Repository struct {
	Reader
	Writer
}

type UseCase interface {
	GetByID(id entities.ID) (*entities.Inventory, error)
	Search(query string) ([]*entities.Inventory, error)
	List() ([]*entities.Inventory, error)
	Create(
		customerID,
		productID entities.ID,
		quantity int,
		unitOfMeasureID entities.ID,
	) entities.ID
	Update(e *entities.Inventory) error
	Delete(id entities.ID) error
}
