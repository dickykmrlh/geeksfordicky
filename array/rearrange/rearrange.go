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

/*
MoveZeroTillTheEnd
Given an array of random numbers, Push all the zero’s of a given array to the end of the array
 */
func MoveZeroTillTheEnd(arr []int) []int{
	nonZeroValueCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			arr[nonZeroValueCount] = arr[i]
			nonZeroValueCount++
		}
	}

	for nonZeroValueCount < len(arr) {
		arr[nonZeroValueCount] = 0
		nonZeroValueCount++
	}

	return arr
}

/*
MinimumSwapRequire
Given an array of n positive integers and a number 'k'.
Find the minimum number of swaps required to bring all the numbers less than or equal to 'k' together.
*/
func MinimumSwapRequire(arr []int, k int) int {
	numberLessThenKNotSubsequentlyTogether := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] <= k && arr[i-1] > k {
			numberLessThenKNotSubsequentlyTogether++
		}
	}
	return numberLessThenKNotSubsequentlyTogether
}

/*
ReOrderArray
Given an array of integers, task is to print the array in the order – the smallest number, the largest number,
2nd smallest number, 2nd largest number, 3rd smallest number, 3rd largest number and so on ...
 */
func ReOrderArray(arr []int) []int{
	sort.Ints(arr)
	var results []int
	i, j := 0, len(arr) - 1
	for  i < j {
		results = append(results, arr[i])
		results = append(results, arr[j])
		i++
		j--
	}

	if i == j {
		results = append(results, arr[i])
	}
	return results
}

/*
DoubleTheFirstElementAndMoveZeroToEnd
For a given array of n integers and assume that ‘0’ is an invalid number and all others as a valid number.
Convert the array in such a way that if both current and next element is valid and both have same value
then double current value and replace the next number with 0. After the modification,
rearrange the array such that all 0’s shifted to the end.
 */
func DoubleTheFirstElementAndMoveZeroToEnd(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}

	i, j := 0, 1
	for j < len(arr) {
		if arr[i] != 0 {
			if arr[i] == arr[j] {
				arr[i] += arr[j]
				arr[j] -= arr[j]
			}
		}
		i++
		j++
	}

	return MoveZeroTillTheEnd(arr)
}
