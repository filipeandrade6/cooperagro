package entities

import "time"

type Inventory struct {
	ID            ID
	User          ID
	Product       ID
	Quantity      int
	UnitOfMeasure ID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewInventory(
	userID,
	productID ID,
	quantity int,
	unitOfMeasureID ID,
) *Inventory {
	return &Inventory{
		ID:            NewID(),
		User:          userID,
		Product:       productID,
		Quantity:      quantity,
		UnitOfMeasure: unitOfMeasureID,
		CreatedAt:     time.Now(),
	}
}
