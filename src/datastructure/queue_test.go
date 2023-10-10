package datastructure

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue(10)
	q.Enqueue(2)
}
