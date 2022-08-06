package customer

import (
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/src/domain/entities"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetByID(id entities.ID) (*entities.Customer, error) {
	c, err := s.repo.GetByID(id)
	if c == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) Search(query string) ([]*entities.Customer, error) {
	customers, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, entities.ErrNotFound
	}

	return customers, nil
}

func (s *Service) List() ([]*entities.Customer, error) {
	customers, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, entities.ErrNotFound
	}

	return customers, nil
}

func (s *Service) Create(
	firstName,
	lastName,
	address,
	phone,
	email string,
	latitude,
	logitude float32,
) (entities.ID, error) {
	c, err := entities.NewCustomer(
		firstName,
		lastName,
		address,
		phone,
		email,
		latitude,
		logitude,
	)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.Create(c)
}

func (s *Service) Update(e *entities.Customer) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.Update(e)
}

func (s *Service) Delete(id entities.ID) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}

	return s.repo.Delete(id)
}
