package datastructure

type maxHeap []any

func (h *maxHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *maxHeap) Push(v any) {
	*h = append(*h, v)
}

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i].(int) < (*h)[j].(int)
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
