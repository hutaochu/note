package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	in := []int{3, 2, 1}
	QuickSort(in, 0, len(in)-1)
	assert.Equal(t, 1, in[0], "in[0] must be 1")
}
