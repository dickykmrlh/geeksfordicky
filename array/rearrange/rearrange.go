package rearrange

import "sort"

/*
SortArrayValueToIndexPosition rearrange array such that the value
match the index position of the array, if the value can't be found for
that particular index, replace it with -1
 */
func SortArrayValueToIndexPosition(arr []int) []int{
	var results []int
	sort.Ints(arr)
	i, j := 0, 0

	for ; i < len(arr); i++ {
		for j < len(arr) {
			if i == arr[j] {
				results = append(results, arr[j])
				j++
				break
			} else if i < arr[j] {
				results = append(results, -1)
				break
			}
			j++
		}
	}

	return results
}
