package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInventory(t *testing.T) {
	i, err := NewInventory(
		NewID(),
		NewID(),
		1,
		NewID(),
	)
	assert.Nil(t, err)
	assert.Equal(t, i.Quantity, 1)
	assert.NotNil(t, i.ID)
}

func TestInventoryValidate(t *testing.T) {
	type test struct {
		userID          ID
		productID       ID
		quantity        int
		unitOfMeasureID ID
		want            error
	}

	tests := []test{
		{
			userID:          NewID(),
			productID:       NewID(),
			quantity:        1,
			unitOfMeasureID: NewID(),
			want:            nil,
		},
		{
			userID:          NewID(),
			productID:       NewID(),
			quantity:        0,
			unitOfMeasureID: NewID(),
			want:            nil,
		},
		{
			userID:          NewID(),
			productID:       NewID(),
			quantity:        -1,
			unitOfMeasureID: NewID(),
			want:            ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := NewInventory(
			tc.userID,
			tc.productID,
			tc.quantity,
			tc.unitOfMeasureID,
		)
		assert.Equal(t, err, tc.want)
	}
}
