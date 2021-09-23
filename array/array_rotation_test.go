package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray(t *testing.T) {
	type args struct {
		arr       []int
		rotations int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "should rotate array correctly",
			args: args{
				arr:       []int{1, 2, 3, 4, 5, 6, 7},
				rotations: 2,
			},
			expected: []int{3, 4, 5, 6, 7, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RotateArrayOneByOne(tt.args.arr, tt.args.rotations)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
