package datastructure_test

import (
	"fmt"
	"github.com/hutaochu/note/datastructure"
	"testing"
)

func TestArray(t *testing.T) {
	array := datastructure.NewArray[int]()
	array.Append(1)
	a, _ := array.Get(0)
	fmt.Println(a)
	array.Append(2)
	a, _ = array.Get(1)
	fmt.Println(a)
}
