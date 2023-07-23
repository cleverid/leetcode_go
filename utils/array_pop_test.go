package utils

import (
	"fmt"
	"lc/utils"
	"testing"
)

func TestPopQueue(t *testing.T) {
	queue := make([]string, 3)
	queue[0] = "a"
	queue[1] = "b"
	queue[2] = "c"
	t.Log(queue)
	x, poped := utils.PopQueue(queue)
	t.Log(*x, poped)
	x2, poped2 := utils.PopQueue(poped)
	t.Log(*x2, poped2)
	x3, poped3 := utils.PopQueue(poped2)
	t.Log(*x3, poped3)
	if x4, poped4 := utils.PopQueue(poped3); x4 == nil {
		fmt.Printf("type %T value %v\n", x4, x4)
		fmt.Printf("slice %v\n", poped4)
	}
}

func TestPopStack(t *testing.T) {
	queue := make([]string, 2)
	queue[0] = "a"
	queue[1] = "b"
	t.Log(queue)
	x, poped := utils.PopStack(queue)
	t.Log(*x, poped)
	x2, poped2 := utils.PopStack(poped)
	t.Log(*x2, poped2)

}
