package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	RoleAdmin    = "admin"
	RoleProducer = "producer"
	RoleBuyer    = "buyer"
)

// User data
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
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User
func NewUser(
	firstName,
	lastName,
	address,
	phone,
	email string,
	latitude,
	longitude float32,
	roles []string,
	password string,
) (*User, error) {
	now := time.Now()

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
		CreatedAt: now,
		UpdatedAt: now,
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}

	c.Password = pwd
	err = c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

// Validate validate User
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

// checkRole verifies if the User roles is valid
func (u *User) checkRoles() bool {
	if len(u.Roles) == 0 {
		return false
	}

	for _, r := range u.Roles {
		if r == RoleAdmin || r == RoleProducer || r == RoleBuyer {
			continue
		}
		return false
	}
	return true
}

func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
