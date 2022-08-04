package usecases

import "github.com/filipeandrade6/cooperagro/src/domain"

type Role int

const (
	Admin Role = iota
	Manager
	Producer
	Buyer
)

type User struct {
	ID       int
	Role     Role
	Customer domain.Customer
}

type UserRepo interface {
	Store(user User)
	FindByID(id int) User
}

type InventoryInteractor struct {
	ProducerRepository      domain.Producer
	BuyerRepository         domain.Buyer
	BaseProductRepository   domain.BaseProduct
	ProductRepository       domain.Product
	UnitOfMeasureRepository domain.UnitOfMeasure
	InventoryRepository     domain.Inventory
	// TODO logger
}

func(i *InventoryInteractor)