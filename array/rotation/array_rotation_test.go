package rotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray(t *testing.T) {
	type args struct {
		arr      []int
		rotation int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "should rotate array correctly",
			args: args{
				arr:      []int{1, 2, 3, 4, 5, 6, 7},
				rotation: 2,
			},
			expected: []int{3, 4, 5, 6, 7, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RotateArrayOneByOne(tt.args.arr, tt.args.rotation)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRotateArrayReverseArray(t *testing.T) {
	type args struct {
		arr      []int
		rotation int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "should rotate array correctly",
			args: args{
				arr:      []int{1, 2, 3, 4, 5, 6, 7},
				rotation: 2,
			},
			expected: []int{3, 4, 5, 6, 7, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RotateArrayReverseArray(tt.args.arr, tt.args.rotation)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFindHowManyTimeArrayHadBeenRotated(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		expected int
	}{
		{
			name: "should return correct rotated counts",
			args: args{
				arr: []int{15, 18, 2, 3, 6, 12},
			},
			expected: 2,
		},
		{
			name: "should return correct rotated counts",
			args: args{
				arr: []int{7, 9, 11, 12, 5},
			},
			expected: 4,
		},
		{
			name: "should return 0, when array is not rotated",
			args: args{
				arr: []int{7, 9, 11, 12, 15},
			},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 actual := FindHowManyTimeArrayHadBeenRotated(tt.args.arr)
			 assert.Equal(t, tt.expected, actual)
		})
	}
}
