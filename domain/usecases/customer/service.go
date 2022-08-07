package customer

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

func (s *Service) GetCustomerByID(id entities.ID) (*entities.Customer, error) {
	c, err := s.repo.GetCustomerByID(id)
	if c == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) SearchCustomer(query string) ([]*entities.Customer, error) {
	customers, err := s.repo.SearchCustomer(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, entities.ErrNotFound
	}

	return customers, nil
}

func (s *Service) ListCustomer() ([]*entities.Customer, error) {
	customers, err := s.repo.ListCustomer()
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, entities.ErrNotFound
	}

	return customers, nil
}

func (s *Service) CreateCustomer(
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

	return s.repo.CreateCustomer(c)
}

func (s *Service) UpdateCustomer(e *entities.Customer) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.UpdateCustomer(e)
}

func (s *Service) DeleteCustomer(id entities.ID) error {
	if _, err := s.GetCustomerByID(id); err != nil {
		return err
	}

	return s.repo.DeleteCustomer(id)
}
