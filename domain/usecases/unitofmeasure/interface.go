package unitofmeasure

import (
	"github.com/filipeandrade6/cooperagro/domain/entities"
)

type Reader interface {
	GetUnitOfMeasureByID(id entities.ID) (*entities.UnitOfMeasure, error)
	SearchUnitOfMeasure(query string) ([]*entities.UnitOfMeasure, error)
	ListUnitOfMeasure() ([]*entities.UnitOfMeasure, error)
}

type Writer interface {
	CreateUnitOfMeasure(e *entities.UnitOfMeasure) (entities.ID, error)
	UpdateUnitOfMeasure(e *entities.UnitOfMeasure) error
	DeleteUnitOfMeasure(id entities.ID) error
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
