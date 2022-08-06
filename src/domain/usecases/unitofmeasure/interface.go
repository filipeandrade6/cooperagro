package unitofmeasure

import (
	"github.com/filipeandrade6/cooperagro/src/domain/entities"
)

type Reader interface {
	GetByID(id entities.ID) (*entities.UnitOfMeasure, error)
	Search(query string) ([]*entities.UnitOfMeasure, error)
	List() ([]*entities.UnitOfMeasure, error)
}

type Writer interface {
	Create(e *entities.UnitOfMeasure) (entities.ID, error)
	Update(e *entities.UnitOfMeasure) error
	Delete(id entities.ID) error
}

type Repository struct {
	Reader
	Writer
}

type UseCase interface {
	GetByID(id entities.ID) (*entities.UnitOfMeasure, error)
	Search(query string) ([]*entities.UnitOfMeasure, error)
	List() ([]*entities.UnitOfMeasure, error)
	Create(name string) (entities.ID, error)
	Update(e *entities.UnitOfMeasure) error
	Delete(id entities.ID) error
}
