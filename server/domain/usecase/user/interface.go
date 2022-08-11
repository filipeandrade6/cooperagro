package user

import "github.com/filipeandrade6/cooperagro/domain/entity"

type Reader interface {
	GetUserByID(id entity.ID) (*entity.User, error)
	SearchUser(query string) ([]*entity.User, error)
	ListUser() ([]*entity.User, error)
}

type Writer interface {
	CreateUser(e *entity.User) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUserByID(id entity.ID) (*entity.User, error)
	SearchUser(query string) ([]*entity.User, error)
	ListUser() ([]*entity.User, error)
	CreateUser(
		firstName,
		lastName,
		address,
		phone,
		email string,
		latitude,
		longitude float32,
		roles []string,
		password string,
	) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}
