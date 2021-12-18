package rotation

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

func isSumOfTwoElementExist(arr []int, sum int) bool {
	r := search.FindPivot(arr, 0, len(arr)-1)
	l := (r + 1)%len(arr)
	for l != r {
		currentSum := arr[l] + arr[r]
		if currentSum == sum {
			return true
		}

		if currentSum > sum {
			r = (len(arr) + r -1) % len(arr)
		} else { // currentSum < sum
			l = (l + 1) % len(arr)
		}
	}

	return false
}
