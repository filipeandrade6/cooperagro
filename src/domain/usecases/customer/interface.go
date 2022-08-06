package customer

import "github.com/filipeandrade6/cooperagro/src/domain/entities"

type Reader interface {
	GetByID(id entities.ID) (*entities.Customer, error)
	Search(query string) ([]*entities.Customer, error)
	List() ([]*entities.Customer, error)
}

type Writer interface {
	Create(e *entities.Customer) (entities.ID, error)
	Update(e *entities.Customer) error
	Delete(id entities.ID) error
}

type Repository struct {
	Reader
	Writer
}

type UseCase interface {
	GetByID(id entities.ID) (*entities.Customer, error)
	Search(query string) ([]*entities.Customer, error)
	List() ([]*entities.Customer, error)
	Create(
		firstName,
		lastName,
		address,
		phone,
		email string,
		latitude,
		longitude float32,
	) (entities.ID, error)
	Update(e *entities.Customer) error
	Delete(id entities.ID) error
}
