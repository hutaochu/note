package datastructure

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   v,
	}
}

func LevelOrder(root *TreeNode) []int {
	queue := list.New()
	queue.PushBack(root)
	var res []int
	for queue.Len() > 0 {
		r := queue.Remove(queue.Front())
		if r == nil {
			break
		}
		node := r.(*TreeNode)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		res = append(res, node.Val)
	}

	return res
}

func PreOrderInsert(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if root.Left == nil {
		root.Left = PreOrderInsert(root.Left, v)
	} else if root.Right == nil {
		root.Right = PreOrderInsert(root.Right, v)
	} else {
		root.Left = PreOrderInsert(root.Left, v)
	}
	return root
}

func PreOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	PreOrderPrint(root.Left)
	PreOrderPrint(root.Right)
}

func DFS(root *TreeNode) {
	stack := list.New()
	stack.PushFront(root)
	for stack.Len() != 0 {
		v, ok := stack.Front().Value.(*TreeNode)
		if ok {
			fmt.Print(v.Val)
			stack.Remove(stack.Front())
		}
		if v.Right != nil {
			stack.PushFront(v.Right)
		}
		if v.Left != nil {
			stack.PushFront(v.Left)
		}
	}
}

func DFSByRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val)
	DFSByRecursion(node.Left)
	DFSByRecursion(node.Right)
}

func BFS(node *TreeNode) {
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() != 0 {
		e := queue.Front()
		v := e.Value.(*TreeNode)
		fmt.Print(v.Val)
		queue.Remove(e)
		if v.Left != nil {
			queue.PushBack(v.Left)
		}
		if v.Right != nil {
			queue.PushBack(v.Right)
		}
	}
}

// func BFSByRecursion(node *TreeNode) {}
