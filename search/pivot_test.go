package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPivot(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "should find pivot on the right half of the array",
			args: args{
				arr:  []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				low:  0,
				high: 8,
			},
			expected: 5,
		},
		{
			name: "should find pivot on the left half of the array, with array length is an odd numbers",
			args: args{
				arr:  []int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7},
				low:  0,
				high: 9,
			},
			expected: 2,
		},
		{
			name: "should not find pivot",
			args: args{
				arr:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				low:  0,
				high: 7,
			},
			expected: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FindPivot(tt.args.arr, tt.args.low, tt.args.high)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
