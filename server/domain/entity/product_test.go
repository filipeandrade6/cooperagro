package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("lima", NewID())
	assert.Nil(t, err)
	assert.Equal(t, p.Name, "lima")
	assert.NotNil(t, p.ID)
}

func TestProductValidate(t *testing.T) {
	type test struct {
		name          string
		baseProductID ID
		want          error
	}

	tests := []test{
		{
			name:          "lima",
			baseProductID: NewID(),
			want:          nil,
		},
		{
			name:          "",
			baseProductID: NewID(),
			want:          ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := NewProduct(tc.name, tc.baseProductID)
		assert.Equal(t, err, tc.want)
	}
}
