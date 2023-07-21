package utils

func FindInArray(arr []int, value int) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}