package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnitOfMeasure(t *testing.T) {
	u, err := NewUnitOfMeasure("kilogram")
	assert.Nil(t, err)
	assert.Equal(t, u.Name, "kilogram")
	assert.NotNil(t, u.ID)
}

func TestUnitOfMeasureValidate(t *testing.T) {
	type test struct {
		name string
		want error
	}

	tests := []test{
		{
			name: "kilogram",
			want: nil,
		},
		{
			name: "",
			want: ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := NewUnitOfMeasure(tc.name)
		assert.Equal(t, err, tc.want)
	}
}
