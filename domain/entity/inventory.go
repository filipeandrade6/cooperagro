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
) *Inventory {
	return &Inventory{
		ID:              NewID(),
		UserID:          userID,
		ProductID:       productID,
		Quantity:        quantity,
		UnitOfMeasureID: unitOfMeasureID,
		CreatedAt:       time.Now(),
	}
}
