package entity_test

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUnitOfMeasure(t *testing.T) {
	bp, err := entity.NewUnitOfMeasure("kilogram")
	assert.Nil(t, err)
	assert.Equal(t, bp.Name, "kilogram")
	assert.NotNil(t, bp.ID)
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
			want: entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewUnitOfMeasure(tc.name)
		assert.Equal(t, err, tc.want)
	}
}
