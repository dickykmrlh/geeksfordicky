package array

func leftRotateByOne(arr []int) {
	firstItem := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = firstItem
}

func RotateArrayOneByOne(arr []int, rotations int) []int {
	for i := 0; i < rotations; i++ {
		leftRotateByOne(arr)
	}
	return arr
}
