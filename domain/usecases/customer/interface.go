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
	GetCustomerByID(id entities.ID) (*entities.Customer, error)
	SearchCustomer(query string) ([]*entities.Customer, error)
	ListCustomer() ([]*entities.Customer, error)
	CreateCustomer(
		firstName,
		lastName,
		address,
		phone,
		email string,
		latitude,
		longitude float32,
	) (entities.ID, error)
	UpdateCustomer(e *entities.Customer) error
	DeleteCustomer(id entities.ID) error
}
