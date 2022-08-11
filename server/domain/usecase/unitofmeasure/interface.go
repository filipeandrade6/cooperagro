package unitofmeasure

import "github.com/filipeandrade6/cooperagro/domain/entity"

type Reader interface {
	GetUnitOfMeasureByID(id entity.ID) (*entity.UnitOfMeasure, error)
	SearchUnitOfMeasure(query string) ([]*entity.UnitOfMeasure, error)
	ListUnitOfMeasure() ([]*entity.UnitOfMeasure, error)
}

type Writer interface {
	CreateUnitOfMeasure(e *entity.UnitOfMeasure) (entity.ID, error)
	UpdateUnitOfMeasure(e *entity.UnitOfMeasure) error
	DeleteUnitOfMeasure(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUnitOfMeasureByID(id entity.ID) (*entity.UnitOfMeasure, error)
	SearchUnitOfMeasure(query string) ([]*entity.UnitOfMeasure, error)
	ListUnitOfMeasure() ([]*entity.UnitOfMeasure, error)
	CreateUnitOfMeasure(name string) (entity.ID, error)
	UpdateUnitOfMeasure(e *entity.UnitOfMeasure) error
	DeleteUnitOfMeasure(id entity.ID) error
}
