package entity

import (
	"time"
)

// Product data
type Product struct {
	ID            ID
	Name          string
	BaseProductID ID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// NewProduct creates a new Product
func NewProduct(name string, baseProductID ID) (*Product, error) {
	p := &Product{
		ID:            NewID(),
		Name:          name,
		BaseProductID: baseProductID,
		CreatedAt:     time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return p, nil
}

// Validate validate Product
func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
