package utils

import (
	"testing"
)

func TestFindInArray(t *testing.T) {
	a := []int{1, 2, 3}
	result := FindInArray(a, 3)
	t.Log(result)
}