package user

import (
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entities"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUserByID(id entities.ID) (*entities.User, error) {
	c, err := s.repo.GetUserByID(id)
	if c == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) SearchUser(query string) ([]*entities.User, error) {
	Users, err := s.repo.SearchUser(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(Users) == 0 {
		return nil, entities.ErrNotFound
	}

	return Users, nil
}

func (s *Service) ListUser() ([]*entities.User, error) {
	Users, err := s.repo.ListUser()
	if err != nil {
		return nil, err
	}
	if len(Users) == 0 {
		return nil, entities.ErrNotFound
	}

	return Users, nil
}

func (s *Service) CreateUser(
	firstName,
	lastName,
	address,
	phone,
	email string,
	latitude,
	logitude float32,
	role string,
) (entities.ID, error) {
	c, err := entities.NewUser(
		firstName,
		lastName,
		address,
		phone,
		email,
		latitude,
		logitude,
		role,
	)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.CreateUser(c)
}

func (s *Service) UpdateUser(e *entities.User) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateUser(e)
}

func (s *Service) DeleteUser(id entities.ID) error {
	if _, err := s.GetUserByID(id); err != nil {
		return err
	}

	return s.repo.DeleteUser(id)
}
