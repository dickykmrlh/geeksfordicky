package rotation

import "geeksfordicky/search"

func leftRotateByOne(arr []int) {
	firstItem := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = firstItem
}

func RotateArrayOneByOne(arr []int, rotation int) []int {
	// in case the rotating factor is
	// greater than array length
	rotation = rotation % len(arr)

	for i := 0; i < rotation; i++ {
		leftRotateByOne(arr)
	}
	return arr
}

func ReverseArray(arr []int, start, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}
func RotateArrayReverseArray(arr []int, rotation int) []int {
	if rotation == 0 {
		return arr
	}

	// in case the rotating factor is
	// greater than array length
	rotation = rotation % len(arr)

	ReverseArray(arr, 0, rotation-1)
	ReverseArray(arr, rotation, len(arr)-1)
	ReverseArray(arr, 0, len(arr)-1)

	return arr
}

func FindHowManyTimeArrayHadBeenRotated(arr []int) int {
	return search.FindPivot(arr, 0, len(arr)-1) + 1
}
