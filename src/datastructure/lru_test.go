package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	assert.Equal(t, -1, c.Get(1), "should get -1")
	c.Put(4, 5)
	c.Put(11, 11)
	assert.Equal(t, 5, c.Get(4), "should get equal")
}
