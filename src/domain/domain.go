package domain

// >> Customer

type CustomerRepo interface {
	Store(customer Customer)
	FindByID(id int) Customer
}

type Customer struct {
	ID   int
	Name string
}

// >> Base product

type BaseProductRepo interface {
	Store(baseProduct BaseProduct)
	FindByID(id int) BaseProduct
}

type BaseProduct struct {
	ID   int
	Name string
}

// >> Product

type ProductRepo interface {
	Store(product Product)
	FindByID(id int) Product
}

type Product struct {
	ID          int
	Name        string
	BaseProduct BaseProduct
}

// >> Unit of measure

type UnitOfMeasureRepo interface {
	Store(unitOfMeasure UnitOfMeasure)
	FindByID(id int) UnitOfMeasure
}

type UnitOfMeasure struct {
	ID   int
	Name string
}

// >> Inventory

type InventoryRepo interface {
	Store(inventory Inventory)
	FindByID(id int) Inventory
}

type Inventory struct {
	ID            int
	Customer      Customer
	Product       Product
	Quantity      int
	UnitOfMeasure UnitOfMeasure
}
