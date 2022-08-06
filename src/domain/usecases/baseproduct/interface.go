package baseproduct

import (
	"github.com/filipeandrade6/cooperagro/src/domain/entities"
)

type Reader interface {
	GetByID(id entities.ID) (*entities.BaseProduct, error)
	Search(query string) ([]*entities.BaseProduct, error)
	List() ([]*entities.BaseProduct, error)
}

type Writer interface {
	Create(e *entities.BaseProduct) (entities.ID, error)
	Update(e *entities.BaseProduct) error
	Delete(id entities.ID) error
}

type Repository struct {
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
