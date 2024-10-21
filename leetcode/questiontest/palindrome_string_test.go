package questiontest

import (
	"testing"

	"github.com/hutaochu/note/leetcode/questionbank"
	"github.com/stretchr/testify/assert"
)

func TestPalindromeString(t *testing.T) {
	s := "abccbaaa"
	r := questionbank.PalindromeString(s)
	assert.Equal(t, "abccba", r, "should be equal")
}
