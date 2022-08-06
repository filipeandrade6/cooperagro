package entities

import "time"

type UnitOfMeasure struct {
	ID        ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUnitOfMeasure(name string) (*UnitOfMeasure, error) {
	u := &UnitOfMeasure{
		ID:        NewID(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return u, nil
}

func (u *UnitOfMeasure) Validate() error {
	if u.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
