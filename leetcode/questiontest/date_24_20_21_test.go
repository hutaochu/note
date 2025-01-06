package questiontest

import (
	"fmt"
	"testing"

	"github.com/hutaochu/note/datastructure"
	"github.com/hutaochu/note/leetcode/questionbank"
)

func TestDistanceK(t *testing.T) {
	root := datastructure.NewTreeNode(3)
	root.Left = datastructure.NewTreeNode(5)
	root.Right = datastructure.NewTreeNode(1)
	root.Left.Left = datastructure.NewTreeNode(6)
	root.Left.Right = datastructure.NewTreeNode(2)
	root.Left.Right.Left = datastructure.NewTreeNode(7)
	root.Left.Right.Right = datastructure.NewTreeNode(4)
	root.Right.Left = datastructure.NewTreeNode(0)
	root.Right.Right = datastructure.NewTreeNode(8)
	res := questionbank.FindDistanceK(root, root.Left, 2)
	fmt.Println(res)
}
