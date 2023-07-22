package etudes

import (
	"testing"
	"fmt"
)

func TestStringIter(t *testing.T) {
	dic := make(map[string]int)
	text := "mama mila ramu"
	for _, c := range text {
		ch := fmt.Sprintf("%c", c)
		fmt.Println(ch)
		dic[ch]++
	}

	t.Log(dic)
}
