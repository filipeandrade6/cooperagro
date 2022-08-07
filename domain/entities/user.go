package entities

import "time"

type User struct {
	ID        ID
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Email     string
	Latitude  float32
	Longitude float32
	Role      []string
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
	role []string,
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
		Role:      role,
		CreatedAt: time.Now(),
	}

	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

func (c *User) Validate() error {
	switch {
	case c.FirstName == "":
		fallthrough
	case c.LastName == "":
		fallthrough
	case c.Address == "":
		fallthrough
	case c.Phone == "":
		fallthrough
	case c.Email == "": // TODO validar e-mail
		fallthrough
	case c.Latitude == 0.0:
		fallthrough
	case c.Longitude == 0.0:
		fallthrough
	case c.Role != "admin" && c.Role != "producer" && c.Role != "buyer":
		return ErrInvalidEntity
	}

	return nil
}
