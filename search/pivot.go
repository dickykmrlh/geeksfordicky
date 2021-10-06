package search

func findPivot(arr []int, low, high int) int {
	if high < low || low > high {
		return -1
	}

	mid := (low + high) / 2

	if mid < high && arr[mid] > arr[mid+1] {
		return mid
	}

	if mid > low && arr[mid] < arr[mid-1] {
		return mid - 1
	}

	if arr[low] > arr[mid] {
		return findPivot(arr, low, mid-1)
	}

	return findPivot(arr, mid+1, high)
}
