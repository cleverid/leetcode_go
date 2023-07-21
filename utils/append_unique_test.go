package utils

import (
	"testing"
)

func TestAppendUnique(t *testing.T) {
	a := []int{1, 2}
	b := []int{1, 2, 3}
	result := AppendUnique(a, b)
	t.Log(result)
}