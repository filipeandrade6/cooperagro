package entity

import (
	"time"
)

type User struct {
	ID        ID
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Email     string
	Latitude  float32
	Longitude float32
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	firstName,
	lastName,
	address,
	phone,
	email string,
	latitude,
	longitude float32,
	roles []string,
) (*User, error) {
	c := &User{
		ID:        NewID(),
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		Phone:     phone,
		Email:     email,
		Latitude:  latitude,
		Longitude: longitude,
		Roles:     roles,
		CreatedAt: time.Now(),
	}

	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

func (u *User) Validate() error {
	switch {
	case u.FirstName == "":
		fallthrough
	case u.LastName == "":
		fallthrough
	case u.Address == "":
		fallthrough
	case u.Phone == "":
		fallthrough
	case u.Email == "": // TODO validar e-mail
		fallthrough
	case u.Latitude == 0.0:
		fallthrough
	case u.Longitude == 0.0:
		fallthrough
	case !u.checkRoles():
		return ErrInvalidEntity
	}

	return nil
}

func (u *User) checkRoles() bool {
	for _, r := range u.Roles {
		if r == "admin" || r == "producer" || r == "buyer" {
			continue
		}
		return false
	}
	return true
}
