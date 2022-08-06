package entities

import "time"

type Customer struct {
	ID        ID
	FirstName string
	LastName  string
	Address   string
	Latitude  float32
	Longitude float32
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(
	firstName,
	lastName,
	address string,
	latitude,
	longitude float32,
	phone,
	email string,
) (*Customer, error) {
	c := &Customer{
		ID:        NewID(),
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
		Phone:     phone,
		Email:     email,
		CreatedAt: time.Now(),
	}

	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

func (p *Customer) Validate() error {
	switch {
	case p.FirstName == "":
		fallthrough
	case p.LastName == "":
		fallthrough
	case p.Address == "":
		fallthrough
	case p.Latitude == 0.0:
		fallthrough
	case p.Longitude == 0.0:
		fallthrough
	case p.Phone == "":
		fallthrough
	case p.Email == "": // TODO validar e-mail
		return ErrInvalidEntity
	}

	return nil
}
