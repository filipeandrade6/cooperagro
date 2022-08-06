package repository

type BaseProductRepository interface {
	GetBaseProductByID()
	SearchBaseProduct()
	ListBaseProduct()
	CreateBaseProduct()
	UpdateBaseProduct()
	DeleteBaseProduct()
}
type CustomerRepository interface {
	GetCustomerByID()
	SearchCustomer()
	ListCustomer()
	CreateCustomer()
	UpdateCustomer()
	DeleteCustomer()
}
type InventoryRepository interface {
	GetInventoryByID()
	SearchInventory()
	ListInventory()
	CreateInventory()
	UpdateInventory()
	DeleteInventory()
}
type ProductRepository interface {
	GetProductByID()
	SearchProduct()
	ListProduct()
	CreateProduct()
	UpdateProduct()
	DeleteProduct()
}
type UnitOfMeasureRepository interface {
	GetUnitOfMesureByID()
	SearchUnitOfMesure()
	ListUnitOfMesure()
	CreateUnitOfMesure()
	UpdateUnitOfMesure()
	DeleteUnitOfMesure()
}

type Repository interface {
	BaseProductRepository
	CustomerRepository
	InventoryRepository
	ProductRepository
	UnitOfMeasureRepository
}
