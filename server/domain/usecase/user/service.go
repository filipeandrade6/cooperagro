package user

import (
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetUserByID(id entity.ID) (*entity.User, error) {
	c, err := s.repo.GetUserByID(id)
	if c == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) SearchUser(query string) ([]*entity.User, error) {
	Users, err := s.repo.SearchUser(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(Users) == 0 {
		return nil, entity.ErrNotFound
	}

	return Users, nil
}

func (s *Service) ListUser() ([]*entity.User, error) {
	Users, err := s.repo.ListUser()
	if err != nil {
		return nil, err
	}
	if len(Users) == 0 {
		return nil, entity.ErrNotFound
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
	roles []string,
	password string,
) (entity.ID, error) {
	c, err := entity.NewUser(
		firstName,
		lastName,
		address,
		phone,
		email,
		latitude,
		logitude,
		roles,
		password,
	)
	if err != nil {
		return entity.NewID(), err
	}

	return s.repo.CreateUser(c)
}

func (s *Service) UpdateUser(e *entity.User) error {
	if err := e.Validate(); err != nil {
		return err
	}

	// TODO: update tem que criptografar a senha novamente
	// TODO: tem que estar nos testes

	e.UpdatedAt = time.Now()

	return s.repo.UpdateUser(e)
}

func (s *Service) DeleteUser(id entity.ID) error {
	if _, err := s.GetUserByID(id); err != nil {
		return err
	}

	return s.repo.DeleteUser(id)
}
