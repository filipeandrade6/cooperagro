package user

import "github.com/filipeandrade6/cooperagro/domain/entities"

type Reader interface {
	GetUserByID(id entities.ID) (*entities.User, error)
	SearchUser(query string) ([]*entities.User, error)
	ListUser() ([]*entities.User, error)
}

type Writer interface {
	CreateUser(e *entities.User) (entities.ID, error)
	UpdateUser(e *entities.User) error
	DeleteUser(id entities.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUserByID(id entities.ID) (*entities.User, error)
	SearchUser(query string) ([]*entities.User, error)
	ListUser() ([]*entities.User, error)
	CreateUser(
		firstName,
		lastName,
		address,
		phone,
		email string,
		latitude,
		longitude float32,
		roles []string,
	) (entities.ID, error)
	UpdateUser(e *entities.User) error
	DeleteUser(id entities.ID) error
}
