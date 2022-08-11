package entity

import "time"

// Inventory data
type Inventory struct {
	ID              ID
	UserID          ID
	ProductID       ID
	Quantity        int
	UnitOfMeasureID ID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// NewInventory creates a new Inventory
func NewInventory(
	userID,
	productID ID,
	quantity int,
	unitOfMeasureID ID,
) (*Inventory, error) {
	now := time.Now()

	i := &Inventory{
		ID:              NewID(),
		UserID:          userID,
		ProductID:       productID,
		Quantity:        quantity,
		UnitOfMeasureID: unitOfMeasureID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	err := i.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return i, nil
}

// Validate validate Inventory
func (i *Inventory) Validate() error {
	if i.Quantity < 0 {
		return ErrInvalidEntity
	}
	return nil
}
