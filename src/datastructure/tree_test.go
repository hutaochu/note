package datastructure

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := NewTreeNode(0)
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	n6 := NewTreeNode(6)

	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Left = n5
	n2.Right = n6
	r := LevelOrder(root)
	fmt.Println(r)
}
