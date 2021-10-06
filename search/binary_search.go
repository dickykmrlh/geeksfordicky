package search

func binarySearch(arr []int, low, high, key int) int {
	if high < low || low > high {
		return -1
	}

	mid := (low + high) / 2
	if key == arr[mid] {
		return mid
	}

	if key > arr[mid] {
		return binarySearch(arr, mid+1, high, key)
	}

	return binarySearch(arr, low, mid-1, key)
}
