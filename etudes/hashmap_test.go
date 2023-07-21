package etudes

import (
	"testing"
)

func TestMapCreate(t *testing.T) {
	a := make(map[string]int)
	a["nastya"] = 31
	t.Log(a)
	value, ok := a["asd"]
	t.Log(value, ok)
}

func TestMapIterate(t * testing.T) {
	a := make(map[string]int)
	a["nastya"] = 31
	a["eugen"] = 35

	t.Log(a)
	for key, value := range a {
		t.Log(key, value)
	}
}
