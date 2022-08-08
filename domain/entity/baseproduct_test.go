package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseProduct(t *testing.T) {
	bp, err := NewBaseProduct("tomate")
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
			want: ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := NewBaseProduct(tc.name)
		assert.Equal(t, err, tc.want)
	}
}
