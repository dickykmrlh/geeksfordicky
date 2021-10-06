package array

import "geeksfordicky/search"

func findKeyInSortedAndRotatedArray(arr []int, key int) int {
	pivot := search.FindPivot(arr, 0, len(arr)-1)

	if pivot == -1 {
		return search.BinarySearch(arr, 0, len(arr)-1, key)
	}

	if arr[pivot] == key {
		return pivot
	}

	if arr[0] < key {
		return search.BinarySearch(arr, 0, pivot-1, key)
	}

	return search.BinarySearch(arr, pivot+1, len(arr)-1, key)
}
