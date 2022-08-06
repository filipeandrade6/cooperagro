package product

import (
	"github.com/filipeandrade6/cooperagro/src/domain/entities"
)

type Reader interface {
	GetByID(id entities.ID) (*entities.Product, error)
	Search(query string) ([]*entities.Product, error)
	List() ([]*entities.Product, error)
}

type Writer interface {
	Create(e *entities.Product) (entities.ID, error)
	Update(e *entities.Product) error
	Delete(id entities.ID) error
}

type Repository struct {
	Reader
	Writer
}

type UseCase interface {
	GetByID(id entities.ID) (*entities.Product, error)
	Search(query string) ([]*entities.Product, error)
	List() ([]*entities.Product, error)
	Create(name string, baseProduct entities.ID) (entities.ID, error)
	Update(e *entities.Product) error
	Delete(id entities.ID) error
}
