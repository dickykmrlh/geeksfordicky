package rotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKeyInSortedAndRotatedArray(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "should return index location when key is found in the right half of the array",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				key: 3,
			},
			expected: 8,
		},
		{
			name: "should return index location when key is found in the left half of the array",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				key: 3,
			},
			expected: 8,
		},
		{
			name: "should return -1 key is not found in the left half of the array",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 10, 1, 2, 3},
				key: 4,
			},
			expected: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findKeyInSortedAndRotatedArray(tt.args.arr, tt.args.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_isSumOfTwoElementExist(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		expected bool
	}{
		{
			name: "should return true, when 2 element with total sum exist",
			args: args{
				arr : []int{11, 15, 6, 8, 9, 10},
				sum: 16,
			},
			expected: true,
		},
		{
			name: "should return true, when 2 element with total sum exist",
			args: args{
				arr : []int{11, 15, 26, 38, 9, 10},
				sum: 35,
			},
			expected: true,
		},
		{
			name: "should return true, when 2 element with total sum exist",
			args: args{
				arr : []int{11, 15, 26, 38, 9, 10},
				sum: 35,
			},
			expected: true,
		},
		{
			name: "should return false, when 2 element with total sum didnt exist",
			args: args{
				arr : []int{11, 15, 26, 38, 9, 10},
				sum: 45,
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isSumOfTwoElementExist(tt.args.arr, tt.args.sum)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
