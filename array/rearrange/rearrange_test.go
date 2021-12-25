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

func TestOddEvenArrangeWithSomeCondition(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		expected []int
	}{
		{
			name: "should correctly arrange the odd and even position",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
			},
			expected: []int{4, 5, 3, 6, 2, 7, 1},
		},
		{
			name: "should correctly arrange the odd and even position take 2",
			args: args{
				arr: []int{1, 2, 1, 4, 5, 6, 8, 8},
			},
			expected: []int{4, 5, 2, 6, 1, 8, 1, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := OddEvenPositionArrangeWithSomeCondition(tt.args.arr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMoveZeroTillTheEnd(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		expected []int
	}{
		{
			name: "should move all 0 element at the very last of array",
			args: args{
				arr: []int{1, 2, 0, 4, 3, 0, 5, 0},
			},
			expected: []int{1, 2, 4, 3, 5, 0, 0, 0},
		},
		{
			name: "should move all 0 element at the very last of array",
			args: args{
				arr: []int{1, 2, 0, 0, 0, 3, 6},
			},
			expected: []int{1, 2, 3, 6, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MoveZeroTillTheEnd(tt.args.arr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
