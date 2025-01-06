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

func TestPreOrderInsert(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)
	PreOrderPrint(root)

	fmt.Println()

	var root1 *TreeNode
	for _, v := range []int{1, 2, 4, 5, 3, 6, 7} {
		if root1 == nil {
			root1 = NewTreeNode(v)
		} else {
			PreOrderInsert(root1, v)
		}
	}
	PreOrderPrint(root1)
}

func TestDFS(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)
	DFS(root)
	fmt.Println("\nDFS use recursion")
	DFSByRecursion(root)

	fmt.Println("\nBFS")
	BFS(root)
}
