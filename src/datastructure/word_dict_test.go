package datastructure_test

import (
	"testing"

	"github.com/hutaochu/note/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestWordDict(t *testing.T) {
	w := datastructure.NewWordDict()
	w.Set("apple")
	w.Set("banana")

	w.Print()
	assert.True(t, w.Exist("apple"))
	assert.False(t, w.Exist("app"))
}
