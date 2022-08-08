package entity_test

import (
	"testing"

	"github.com/filipeandrade6/cooperagro/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseProduct(t *testing.T) {
	bp, err := entity.NewBaseProduct("tomate")
	assert.Nil(t, err)
	assert.Equal(t, bp.Name, "tomate")
	assert.NotNil(t, bp.ID)
}

func TestBaseProductValidate(t *testing.T) {
	type test struct {
		name string
		want error
	}

	tests := []test{
		{
			name: "tomate",
			want: nil,
		},
		{
			name: "",
			want: entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewBaseProduct(tc.name)
		assert.Equal(t, err, tc.want)
	}
}
