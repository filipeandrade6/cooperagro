package mock

import (
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type MockUserService struct {
	validID entity.ID
}

func NewMockUserService(validID entity.ID) MockUserService {
	return MockUserService{validID: validID}
}

func (m MockUserService) GetUserByID(id entity.ID) (*entity.User, error) {
	if id == m.validID {
		return &entity.User{
			ID:        id,
			FirstName: "Filipe",
			LastName:  "Andrade",
			Address:   "Main street",
			Phone:     "5561555554444",
			Email:     "filipe@email.com",
			Latitude:  -12.123456,
			Longitude: -12.123456,
			Roles:     []string{"admin", "producer"},
			Password:  "$2a$10$nTxInXic3WCz14l64ycdx.78LJxBNGcw4/yT4LkdD9WZmFwAy/.pW",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	return nil, entity.ErrNotFound
}

func (m MockUserService) SearchUser(query string) ([]*entity.User, error) {
	return nil, nil
}

func (m MockUserService) ListUser() ([]*entity.User, error) {
	return nil, nil
}

func (m MockUserService) CreateUser(
	firstName,
	lastName,
	address,
	phone,
	email string,
	latitude,
	longitude float32,
	roles []string,
	password string,
) (entity.ID, error) {
	return entity.NewID(), nil
}

func (m MockUserService) UpdateUser(e *entity.User) error {
	return nil
}

func (m MockUserService) DeleteUser(id entity.ID) error {
	return nil
}
