package rearrange

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
