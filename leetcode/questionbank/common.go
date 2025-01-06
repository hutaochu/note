package questionbank

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Var  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l != nil && l.Var != -1 {
		fmt.Print(l.Var, " ")
	}
	if l.Next != nil {
		l.Next.Print()
	}
}

type minHeapForList struct {
	data []*ListNode
}

func (h *minHeapForList) Print() {
	fmt.Println()
	for _, v := range h.data {
		fmt.Print(v.Var, " ")
	}
	fmt.Println()
}

func (h *minHeapForList) Less(i, j int) bool {
	return h.data[i].Var < h.data[j].Var
}

func (h *minHeapForList) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *minHeapForList) Len() int {
	return len(h.data)
}

func (h *minHeapForList) Pop() any {
	tail := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return tail
}

func (h *minHeapForList) Push(item any) {
	h.data = append(h.data, item.(*ListNode))
}

// MergeList
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func MergeList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	p := head
	p1, p2 := l1, l2
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p.Next = p2
			break
		}
		if p2 == nil {
			p.Next = p1
			break
		}
		if p1.Var < p2.Var {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	return head
}

// PartitionList
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
func PartitionList(l *ListNode, target int) *ListNode {
	p1, p2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	l1, l2 := p1, p2
	cur := l
	for cur != nil {
		if cur.Var < target {
			l1.Next = cur
			l1 = l1.Next
		} else {
			l2.Next = cur
			l2 = l2.Next
		}
		cur = cur.Next
	}
	l2.Next = nil
	l1.Next = p2
	return p1
}

// MergeKList
// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func MergeKList(in []*ListNode) *ListNode {
	h := &minHeapForList{}
	heap.Init(h)
	for _, v := range in {
		temp := v
		for temp != nil {
			heap.Push(h, temp)
			temp = temp.Next
		}
	}
	res := &ListNode{-1, nil}
	p := res

	h.Print()
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		v.Next = nil
		p.Next = v
		p = p.Next
		h.Print()
	}
	return res
}

// FindKthLastInList
// 单链表的倒数第 k 个节点
func FindKthLastInList(list *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}
	p := list
	kl := list
	counter := 1
	for p != nil {
		p = p.Next
		if counter > k {
			kl = kl.Next
		}
		counter++
	}
	kl.Next = nil
	return kl
}

// FindMidNode
//
// 查找单链表的中点
func FindMidNode(l *ListNode) *ListNode {
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// HasCycle
//
// 检查链表中是否存在环
func HasCycle(l *ListNode) bool {
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// FindCycleStartNode
//
// 判断是否有环，若有，则返回环的起点
func FindCycleStartNode(l *ListNode) *ListNode {
	slow, fast := l, l
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if hasCycle {
		fast = l
		for fast != slow {
			slow = slow.Next
			fast = fast.Next
		}
		return fast
	}
	return nil
}

// RemoveDuplicates
//
// 删除非严格递增数组中的重复元素
func RemoveDuplicates(nums []int) int {
	p1, p2 := 0, 0
	for p2 <= len(nums)-1 {
		if nums[p1] != nums[p2] {
			p1++
			nums[p1] = nums[p2]
		}
		p2++
	}
	return p1 + 1
}

func FindKSum(a1, a2 []int, k int) []int {
	// get top k
	h := &myHeap{}
	heap.Init(h)
	for i := len(a1) - 1; i >= 0; i-- {
		for j := len(a2) - 1; j >= 0; j-- {
			heap.Push(h, a1[i]+a2[j])
		}
	}
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(int))
	}

	return result
}

type myHeap []int

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) Pop() interface{} {
	tail := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tail
}

type Item struct {
	sum int
	i   int
	j   int
}

type MaxHeap []Item

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i].sum > (*h)[j].sum
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i].sum, (*h)[j].sum = (*h)[j].sum, (*h)[i].sum
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Push(v any) {
	*h = append(*h, v.(Item))
}

func (h *MaxHeap) Pop() any {
	tail := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return tail
}

func FindKLargestSums(A, B []int, k int) []int {
	var results []int
	if len(A) == 0 || len(B) == 0 || k <= 0 {
		return results
	}
	m, n := len(A)-1, len(B)-1
	maxHeap := &MaxHeap{
		{
			sum: A[m] + B[n],
			i:   m,
			j:   n,
		},
	}
	heap.Init(maxHeap)
	top := (*maxHeap)[0]
	visited := make(map[[2]int]bool)
	visited[[2]int{m, n}] = true
	for maxHeap.Len() > 0 && k > 0 {
		top = heap.Pop(maxHeap).(Item)
		results = append(results, top.sum)
		k--

		if top.i > 0 && !visited[[2]int{top.i - 1, top.j}] {
			heap.Push(maxHeap, Item{
				sum: A[top.i-1] + B[top.j],
				i:   top.i - 1,
				j:   top.j,
			})
			visited[[2]int{top.i - 1, top.j}] = true
		}
		if top.j > 0 && !visited[[2]int{top.i, top.j - 1}] {
			heap.Push(maxHeap, Item{
				sum: A[top.i] + B[top.j-1],
				i:   top.i,
				j:   top.j - 1,
			})
			visited[[2]int{top.i, top.j - 1}] = true
		}

	}
	return results
}
