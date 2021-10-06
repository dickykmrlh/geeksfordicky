package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
		key  int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "should find correct index location for key located on left side of the array",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				low:  0,
				high: 8,
				key:  3,
			},
			expected: 2,
		},
		{
			name: "should find correct index location for key located on right side of the array, and array length is odd",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				low:  0,
				high: 9,
				key:  8,
			},
			expected: 7,
		},
		{
			name: "should return -1, when it cant find the key in the array",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				low:  0,
				high: 9,
				key:  11,
			},
			expected: -1,
		},
		{
			name: "should return -1, when it cant find the key in the array",
			args: args{
				arr:  []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
				low:  0,
				high: 8,
				key:  1,
			},
			expected: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := binarySearch(tt.args.arr, tt.args.low, tt.args.high, tt.args.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
