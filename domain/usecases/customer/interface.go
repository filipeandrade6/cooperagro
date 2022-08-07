package customer

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Reader interface {
	GetCustomerByID(id entities.ID) (*entities.Customer, error)
	SearchCustomer(query string) ([]*entities.Customer, error)
	ListCustomer() ([]*entities.Customer, error)
}

type Writer interface {
	CreateCustomer(e *entities.Customer) (entities.ID, error)
	UpdateCustomer(e *entities.Customer) error
	DeleteCustomer(id entities.ID) error
}

type Repository interface {
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
