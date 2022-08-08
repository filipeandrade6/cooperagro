package entity_test

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := entity.NewProduct("lima", entity.NewID())
	assert.Nil(t, err)
	assert.Equal(t, p.Name, "lima")
	assert.NotNil(t, p.ID)
}

func TestProductValidate(t *testing.T) {
	type test struct {
		name          string
		baseProductID entity.ID
		want          error
	}

	tests := []test{
		{
			name:          "lima",
			baseProductID: entity.NewID(),
			want:          nil,
		},
		{
			name:          "",
			baseProductID: entity.NewID(),
			want:          entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewProduct(tc.name, tc.baseProductID)
		assert.Equal(t, err, tc.want)
	}
}
