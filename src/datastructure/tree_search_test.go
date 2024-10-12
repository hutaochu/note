package datastructure

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func printTree(node *TreeSearchNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	// 打印当前节点
	var branch string
	if isLeft {
		branch = "├──"
	} else {
		branch = "└──"
	}
	fmt.Println(prefix + branch + strconv.Itoa(node.Key))

	// 构建前缀并递归打印左右子树
	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}
	printTree(node.Left, newPrefix, true)
	printTree(node.Right, newPrefix, false)
}

func printTreeVertically(node *TreeSearchNode, space int, level int) {
	if node == nil {
		return
	}

	// 每个节点增加的间距
	increment := 5

	// 先打印右子树，增加间距
	printTreeVertically(node.Right, space+increment, level+1)

	// 打印当前节点
	fmt.Println()
	for i := 0; i < space; i++ {
		fmt.Print(" ") // 打印空格，使节点对齐
	}
	fmt.Println(node.Key) // 打印当前节点值

	// 打印左子树
	printTreeVertically(node.Left, space+increment, level+1)
}

func getRandSlice(count, max int) []int {
	list := make([]int, count)
	for i := 0; i < count; i++ {
		list[i] = rand.Intn(max)
	}
	return list
}

func TestTreeSearch(t *testing.T) {
	max := 100
	node := NewTreeSearch(max / 2)
	list := getRandSlice(5, 100)
	fmt.Println(list)
	for _, v := range list {
		Insert(&node, v)
		fmt.Println("insert: ", v)
		printTreeVertically(&node, 0, 0)
		fmt.Println("--------------------")
	}

	// printTreeVertically(&node, 0, 0)
}
