package entity

import "time"

type BaseProduct struct {
	ID        ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

func (bp *BaseProduct) Validate() error {
	if bp.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
