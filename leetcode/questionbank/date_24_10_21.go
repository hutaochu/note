package questionbank

import (
	"container/list"
	"github.com/hutaochu/note/datastructure"
)

// FindDistanceK godoc
// the random medium question number is 863
// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，
//
// 和一个整数值 k ，返回到目标结点 target 距离为 k 的所有结点的值的数组。
//
// 答案可以以 任何顺序 返回。
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
//
// 输出：[7,4,1]
//
// 解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1
func FindDistanceK(root *datastructure.TreeNode, target *datastructure.TreeNode, k int) []int {
	parentMap := make(map[*datastructure.TreeNode]*datastructure.TreeNode)
	// DFT, get all relations of node and their parent
	var dfs func(node *datastructure.TreeNode)
	dfs = func(node *datastructure.TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parentMap[node.Left] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	// BFS, get all k distance node
	checkedNode := make(map[*datastructure.TreeNode]struct{})
	queue := list.New()
	distance := 0
	queue.PushBack(target)
	for queue.Len() != 0 && distance <= k {
		// find result
		if distance == k {
			var res []int
			for queue.Len() != 0 {
				e := queue.Front()
				res = append(res, e.Value.(*datastructure.TreeNode).Val)
				queue.Remove(e)
			}
			return res
		}

		for i := queue.Len(); i > 0; i-- {
			e := queue.Front()
			node := e.Value.(*datastructure.TreeNode)
			if _, ok := checkedNode[node]; !ok {
				checkedNode[node] = struct{}{}
				if p, ok := parentMap[node]; ok && p != nil {
					if _, ok := checkedNode[p]; !ok {
						queue.PushBack(p)
					}

				}
				if node.Left != nil {
					if _, ok := checkedNode[node.Left]; !ok {
						queue.PushBack(node.Left)
					}
				}
				if node.Right != nil {
					if _, ok := checkedNode[node.Right]; !ok {
						queue.PushBack(node.Right)
					}
				}
			}

			queue.Remove(e)
			//fmt.Println(node, e.Value.(*datastructure.TreeNode).Val)
		}

		distance++
	}

	return []int{}
}
