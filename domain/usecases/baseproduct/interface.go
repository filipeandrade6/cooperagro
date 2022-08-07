package baseproduct

import (
	"github.com/filipeandrade6/cooperagro/domain/entities"
)

type Reader interface {
	GetBaseProductByID(id entities.ID) (*entities.BaseProduct, error)
	SearchBaseProduct(query string) ([]*entities.BaseProduct, error)
	ListBaseProduct() ([]*entities.BaseProduct, error)
}

type Writer interface {
	CreateBaseProduct(e *entities.BaseProduct) (entities.ID, error)
	UpdateBaseProduct(e *entities.BaseProduct) error
	DeleteBaseProduct(id entities.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetByID(id entities.ID) (*entities.BaseProduct, error)
	Search(query string) ([]*entities.BaseProduct, error)
	List() ([]*entities.BaseProduct, error)
	Create(name string) (entities.ID, error)
	Update(e *entities.BaseProduct) error
	Delete(id entities.ID) error
}
