package entity

import "time"

// BaseProduct data
type BaseProduct struct {
	ID        ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewBaseProduct creates a new BaseProduct
func NewBaseProduct(name string) (*BaseProduct, error) {
	bp := &BaseProduct{
		ID:        NewID(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	err := bp.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return bp, nil
}

// Validate validate BaseProduct
func (bp *BaseProduct) Validate() error {
	if bp.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
