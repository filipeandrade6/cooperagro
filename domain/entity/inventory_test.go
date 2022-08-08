package entity_test

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewInventory(t *testing.T) {
	i, err := entity.NewInventory(
		entity.NewID(),
		entity.NewID(),
		1,
		entity.NewID(),
	)
	assert.Nil(t, err)
	assert.Equal(t, i.Quantity, 1)
	assert.NotNil(t, i.ID)
}

func TestInventoryValidate(t *testing.T) {
	type test struct {
		userID          entity.ID
		productID       entity.ID
		quantity        int
		unitOfMeasureID entity.ID
		want            error
	}

	tests := []test{
		{
			userID:          entity.NewID(),
			productID:       entity.NewID(),
			quantity:        1,
			unitOfMeasureID: entity.NewID(),
			want:            nil,
		},
		{
			userID:          entity.NewID(),
			productID:       entity.NewID(),
			quantity:        0,
			unitOfMeasureID: entity.NewID(),
			want:            nil,
		},
		{
			userID:          entity.NewID(),
			productID:       entity.NewID(),
			quantity:        -1,
			unitOfMeasureID: entity.NewID(),
			want:            entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewInventory(
			tc.userID,
			tc.productID,
			tc.quantity,
			tc.unitOfMeasureID,
		)
		assert.Equal(t, err, tc.want)
	}
}
