package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	InsertSort(arr)
	assert.Equal(t, 1, arr[0])
	fmt.Println(arr)
}
