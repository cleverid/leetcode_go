package utils

func AppendUnique(arr1 []int, arr2 []int) []int {
	for _, item := range arr2 {
		if !FindInArray(arr1, item) {
			arr1 = append(arr1, item)
		}
	}
	return arr1
}