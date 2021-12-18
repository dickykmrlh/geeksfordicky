package rearrange

import "sort"

/*
SortArrayValueToIndexPosition rearrange array such that the value
match the index position of the array, if the value can't be found for
that particular index, replace it with -1
 */
func SortArrayValueToIndexPosition(arr []int) []int{
	m := make(map[int]int)

	for _, v := range arr {
		m[v] = v
	}

	for i, _ := range arr {
		val, ok := m[i]
		if ok {
			arr[i] = val
		} else {
			arr[i] = -1
		}
	}

	return arr
}

/*
OddEvenPositionArrangeWithSomeCondition
Rearrange array such that arr[i] >= arr[j] if i is even and arr[i]<=arr[j] if i is odd and j < i
Given an array of n elements. Our task is to write a program to rearrange the array such that
elements at even positions are greater than all elements before it and elements at odd positions are
less than all elements before it.
 */
func OddEvenPositionArrangeWithSomeCondition(arr []int) []int {
	totalEven := len(arr) / 2
	totalOdd := len(arr) - totalEven
	var tempArr []int
	for _, v := range arr {
		tempArr = append(tempArr, v)
	}

	sort.Ints(tempArr)

	// fill odd position
	oddPosition := totalOdd - 1
	for i := 0; i < len(arr); i+=2 {
		arr[i] = tempArr[oddPosition]
		oddPosition--
	}

	// fill even position
	evenPosition := totalOdd
	for i := 1; i < len(arr); i+=2 {
		arr[i] = tempArr[evenPosition]
		evenPosition++
	}

	return arr
}
