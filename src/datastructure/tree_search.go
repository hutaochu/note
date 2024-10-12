package datastructure

type TreeSearchNode struct {
	P     *TreeSearchNode
	Key   int
	Left  *TreeSearchNode
	Right *TreeSearchNode
}

func NewTreeSearch(v int) TreeSearchNode {
	return TreeSearchNode{
		Key: v,
	}
}

func Insert(root *TreeSearchNode, v int) {
	pre := root
	next := root
	for next != nil {
		pre = next
		if v <= next.Key {
			next = next.Left
		} else {
			next = next.Right
		}
	}
	newNode := TreeSearchNode{
		P:   pre,
		Key: v,
	}
	if v <= pre.Key {
		pre.Left = &newNode
	} else {
		pre.Right = &newNode
	}
}

func Next(node *TreeSearchNode) *TreeSearchNode {
	next := node
	if next.Right == nil {
		return next
	}

	for next != nil {
		if next.Left == nil {
			return next
		}
		next = next.Left
	}
	return nil
}

func Delete(root *TreeSearchNode, v int) {

}
