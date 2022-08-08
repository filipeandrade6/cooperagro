package product

import "github.com/filipeandrade6/cooperagro/domain/entity"

type Reader interface {
	GetProductByID(id entity.ID) (*entity.Product, error)
	SearchProduct(query string) ([]*entity.Product, error)
	ListProduct() ([]*entity.Product, error)
}

type Writer interface {
	CreateProduct(e *entity.Product) (entity.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetProductByID(id entity.ID) (*entity.Product, error)
	SearchProduct(query string) ([]*entity.Product, error)
	ListProduct() ([]*entity.Product, error)
	CreateProduct(name string, baseProduct entity.ID) (entity.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id entity.ID) error
}
