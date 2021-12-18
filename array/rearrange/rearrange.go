package rearrange

/*
SortArrayValueToIndexPosition rearrange array such that the value
match the index position of the array, if the value can't be found for
that particular index, replace it with -1
 */
func SortArrayValueToIndexPosition(arr []int) []int{
	var results []int
	for index, _ := range arr {
		found := false
		updatedValue := 0
		for _, value := range arr {
			if index == value {
				found = true
				updatedValue = value
			}
		}

		if found {
			results = append(results, updatedValue)
		} else {
			results = append(results, -1)
		}
	}

	return results
}
