package domain

import "time"

type UnitOfMeasure struct {
	ID   int
	Name string
}

type BaseProduct struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Product struct {
	ID              int
	Name            string
	BaseProductID   int
	UnitOfMeasureID int
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type Role struct {
	ID   int
	Role string
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	Latitude  float32
	Longitude float32
	RoleID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
