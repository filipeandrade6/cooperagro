package entity

import "time"

// UnitOfMeasure data
type UnitOfMeasure struct {
	ID        ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUnitOfMeasure creates a new UnitOfMeasure
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

// Validate validate UnitOfMeasure
func (u *UnitOfMeasure) Validate() error {
	if u.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
