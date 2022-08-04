package domain

type Producer struct {
	ID   int
	Name string
}

type Buyer struct {
	ID   int
	Name string
}

type BaseProduct struct {
	ID   int
	Name string
}

type Product struct {
	ID          int
	Name        string
	BaseProduct BaseProduct
}

type UnitOfMeasure struct {
	ID   int
	Name string
}

type Inventory struct {
	ID            int
	Producer      Producer
	Product       Product
	Quantity      int
	UnitOfMeasure UnitOfMeasure
}
