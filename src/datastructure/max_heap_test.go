package datastructure

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := NewMyMaxHeap()
	h.Push(2)
	h.Push(5)
	h.Push(3)
	h.Push(10)
	fmt.Println(h)
}
