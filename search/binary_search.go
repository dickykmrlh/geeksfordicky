package search

func BinarySearch(arr []int, low, high, key int) int {
	if high < low || low > high {
		return -1
	}

	mid := (low + high) / 2
	if key == arr[mid] {
		return mid
	}

	if key > arr[mid] {
		return BinarySearch(arr, mid+1, high, key)
	}

	return BinarySearch(arr, low, mid-1, key)
}
