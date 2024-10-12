package datastructure

import "container/list"

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
