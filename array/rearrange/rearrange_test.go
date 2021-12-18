package rearrange

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRearrangeMatchIndexPosition(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		expected []int
	}{
		{
			name: "Should rearrange array correctly, with -1 in missing value on that index position",
			args: args{
				arr: []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1},
			},
			expected: []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9},
		},
		{
			name: "Should rearrange array correctly, with -1 in missing value on that index position take 2",
			args: args{
				arr: []int{19, 7, 0, 3, 18, 15, 12, 6, 1, 8, 11, 10, 9, 5, 13, 16, 2, 14, 17, 4},
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SortArrayValueToIndexPosition(tt.args.arr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}