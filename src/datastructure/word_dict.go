package datastructure

import "fmt"

/*
可以引⼊的库和版本相关请参考 “环境说明”
Please refer to the "Environmental Notes" for the libraries and versions that can be introduced.

api 1: give you a word                                 apple
api 2: have you seen a specific word?                  apple? true    banana? false
api 3: have you seen a word with a specific prefix?    app? true      abc? false
*/

type WordNode struct {
	Element map[rune]*WordNode
}

type WordDict struct {
	Node *WordNode
}

func NewWordDict() *WordDict {
	return &WordDict{
		Node: &WordNode{
			Element: make(map[rune]*WordNode),
		},
	}
}

func (n *WordNode) Print() {
	for k, v := range n.Element {
		fmt.Print(string(k))
		if len(v.Element) != 0 {
			v.Print()
		} else {
			fmt.Println()
		}
	}
}

func (n *WordNode) Set(w rune) *WordNode {
	if next, ok := n.Element[w]; ok {
		return next
	}
	n.Element[w] = &WordNode{
		Element: make(map[rune]*WordNode),
	}
	return n.Element[w]
}

func (w *WordDict) Set(word string) {
	curNode := w.Node
	for _, r := range word {
		curNode = curNode.Set(r)
	}
}

func (w *WordDict) Exist(word string) bool {
	curNode := w.Node
	for _, r := range word {
		if _, ok := curNode.Element[r]; ok {
			curNode = curNode.Element[r]
		} else {
			// not found
			return false
		}
	}
	// 如果有子节点，说明不是完整的单词
	if len(curNode.Element) != 0 {
		return false
	}
	return true
}

func (w *WordDict) ExistWithPrefix(prefix string) bool {
	curNode := w.Node
	for _, r := range prefix {
		if _, ok := curNode.Element[r]; ok {
			curNode = curNode.Element[r]
		} else {
			// not found
			return false
		}
	}

	return true
}

func (w *WordDict) Print() {
	w.Node.Print()
}
