package domain

// >> Customer

type Customer struct {
	ID   int
	Name string
}

type CustomerRepo interface {
	Store(customer Customer)
	FindByID(id int) Customer
}

// >> Base product

type BaseProduct struct {
	ID   int
	Name string
}

type BaseProductRepo interface {
	Store(baseProduct BaseProduct)
	FindByID(id int) BaseProduct
}

// >> Product

type Product struct {
	ID          int
	Name        string
	BaseProduct BaseProduct
}

type ProductRepo interface {
	Store(product Product)
	FindByID(id int) Product
}

// >> Unit of measure

type UnitOfMeasure struct {
	ID   int
	Name string
}

type UnitOfMeasureRepo interface {
	Store(unitOfMeasure UnitOfMeasure)
	FindByID(id int) UnitOfMeasure
}

// >> Inventory

type Inventory struct {
	ID            int
	Customer      Customer
	Product       Product
	Quantity      int
	UnitOfMeasure UnitOfMeasure
}

type InventoryRepo interface {
	Store(inventory Inventory)
	FindByID(id int) Inventory
}
