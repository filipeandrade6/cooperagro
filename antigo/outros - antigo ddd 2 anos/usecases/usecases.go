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
	UserRepo          UserRepo
	BaseProductRepo   domain.BaseProductRepo
	ProductRepo       domain.ProductRepo
	UnitOfMeasureRepo domain.UnitOfMeasureRepo
	InventoryRepo     domain.InventoryRepo
	// TODO logger
}

func (i InventoryInteractor) Create(userID, productID, unitOfMeasereID, quantity int) (int, error) {
	return 0, nil
}

func (i InventoryInteractor) Read(inventoryID int) (domain.Inventory, error) {
	return domain.Inventory{}, nil
}

func (i InventoryInteractor) Update(userID int, inventory domain.Inventory) error {
	return nil
}

func (i InventoryInteractor) Delete(inventoryID int) error {
	return nil
}

type ManagerInventoryInteractor struct {
	InventoryInteractor InventoryInteractor
}
