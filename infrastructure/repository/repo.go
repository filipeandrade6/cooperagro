package repository

import "github.com/filipeandrade6/cooperagro/domain/entities"

type BaseProductRepository interface {
	GetBaseProductByID(id entities.ID) (*entities.BaseProduct, error)
	SearchBaseProduct(query string) ([]*entities.BaseProduct, error)
	ListBaseProduct() ([]*entities.BaseProduct, error)
	CreateBaseProduct(e *entities.BaseProduct) (entities.ID, error)
	UpdateBaseProduct(e *entities.BaseProduct) error
	DeleteBaseProduct(id entities.ID) error
}

type CustomerRepository interface {
	GetCustomerByID(id entities.ID) (*entities.Customer, error)
	SearchCustomer(query string) ([]*entities.Customer, error)
	ListCustomer() ([]*entities.Customer, error)
	CreateCustomer(e *entities.Customer) (entities.ID, error)
	UpdateCustomer(e *entities.Customer) error
	DeleteCustomer(id entities.ID) error
}

type InventoryRepository interface {
	GetInventoryByID()
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
