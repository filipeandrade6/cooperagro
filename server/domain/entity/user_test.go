package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	bp, err := NewUser(
		"Filipe",
		"Andrade",
		"Main street",
		"5561999994444",
		"filipe.engenharia42@gmail.com",
		-12.123456,
		-12.123456,
		[]string{"admin", "producer", "buyer"},
		"admin",
	)

	assert.Nil(t, err)
	assert.Equal(t, bp.FirstName, "Filipe")
	assert.NotNil(t, bp.ID)
}

func TestUserValidate(t *testing.T) {
	type test struct {
		firstName string
		lastName  string
		address   string
		phone     string
		email     string
		latitude  float32
		longitude float32
		roles     []string
		password  string
		want      error
	}

	tests := []test{
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      nil,
		},
		{
			firstName: "",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  0.0,
			longitude: -12.123456,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: 0.0,
			roles:     []string{"producer", "buyer"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     nil,
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"research"},
			password:  "admin",
			want:      ErrInvalidEntity,
		},
		{
			firstName: "Filipe",
			lastName:  "Andrade",
			address:   "Main street",
			phone:     "5561555554444",
			email:     "filipe@email.com",
			latitude:  -12.123456,
			longitude: -12.123456,
			roles:     []string{"research"},
			password:  "",
			want:      ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := NewUser(
			tc.firstName,
			tc.lastName,
			tc.address,
			tc.phone,
			tc.email,
			tc.latitude,
			tc.longitude,
			tc.roles,
			tc.password,
		)
		assert.Equal(t, err, tc.want)
	}
}
