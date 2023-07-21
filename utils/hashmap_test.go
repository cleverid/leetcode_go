package utils

import (
	. "lc/utils"
	"testing"
)

func TestMapKeys(t *testing.T) {
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2

	keys := MapKeys(a)
	t.Log(a)
	t.Log(keys)
}

func TestMapValues(t *testing.T) {
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2

	values := MapValues(a)
	t.Log(a)
	t.Log(values)
}
