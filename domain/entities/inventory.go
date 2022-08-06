package entities

import "time"

type Inventory struct {
	ID            ID
	Customer      ID
	Product       ID
	Quantity      int
	UnitOfMeasure ID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewInventory(
	customerID,
	productID ID,
	quantity int,
	unitOfMeasureID ID,
) *Inventory {
	return &Inventory{
		ID:            NewID(),
		Customer:      customerID,
		Product:       productID,
		Quantity:      quantity,
		UnitOfMeasure: unitOfMeasureID,
		CreatedAt:     time.Now(),
	}
}
