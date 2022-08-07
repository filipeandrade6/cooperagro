package product

import (
	"github.com/filipeandrade6/cooperagro/domain/entities"
)

type Reader interface {
	GetProductByID(id entities.ID) (*entities.Product, error)
	SearchProduct(query string) ([]*entities.Product, error)
	ListProduct() ([]*entities.Product, error)
}

type Writer interface {
	CreateProduct(e *entities.Product) (entities.ID, error)
	UpdateProduct(e *entities.Product) error
	DeleteProduct(id entities.ID) error
}

type Repository interface {
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
