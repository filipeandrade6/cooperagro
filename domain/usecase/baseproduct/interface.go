package baseproduct

import "github.com/filipeandrade6/cooperagro/domain/entity"

type Reader interface {
	GetBaseProductByID(id entity.ID) (*entity.BaseProduct, error)
	SearchBaseProduct(query string) ([]*entity.BaseProduct, error)
	ListBaseProduct() ([]*entity.BaseProduct, error)
}

type Writer interface {
	CreateBaseProduct(e *entity.BaseProduct) (entity.ID, error)
	UpdateBaseProduct(e *entity.BaseProduct) error
	DeleteBaseProduct(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetBaseProductByID(id entity.ID) (*entity.BaseProduct, error)
	SearchBaseProduct(query string) ([]*entity.BaseProduct, error)
	ListBaseProduct() ([]*entity.BaseProduct, error)
	CreateBaseProduct(name string) (entity.ID, error)
	UpdateBaseProduct(e *entity.BaseProduct) error
	DeleteBaseProduct(id entity.ID) error
}
