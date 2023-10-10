package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	v, _ := stack.Pop()
	assert.Equal(t, v, 3)
	stack.Pop()
	stack.Pop()
	_, ok := stack.Pop()
	assert.Equal(t, ok, false)
}
