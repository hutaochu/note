package datastructure_test

import (
	"testing"

	"github.com/hutaochu/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	a := datastructure.NewArray(5)

	a.Insert(0, 1)
	a.Insert(1, 2)
	a.Insert(2, 3)
	a.Insert(3, 4)
	a.Insert(4, 5)
	assert.Equal(t, 3, a.Get(2))
	assert.Equal(t, 5, a.Capacity())
	assert.Equal(t, 5, a.Length())

	a.Remove(3)
	assert.Equal(t, 5, a.Get(3))
}
