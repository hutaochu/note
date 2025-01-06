package datastructure

type MyMaxHeap struct {
	data []int
}

func NewMyMaxHeap() *MyMaxHeap {
	return &MyMaxHeap{
		data: make([]int, 0),
	}
}

func (h *MyMaxHeap) left(i int) int {
	return 2*i + 1
}

func (h *MyMaxHeap) right(i int) int {
	return 2*i + 2
}

func (h *MyMaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MyMaxHeap) Push(v int) {
	h.data = append(h.data, v)
	h.up(h.len() - 1)
}

func (h *MyMaxHeap) Pop() int {
	h.swap(0, h.len()-1)
	last := h.data[h.len()-1]
	h.data = h.data[:h.len()-1]
	h.down(0)
	return last
}

// up 从底到顶做“堆化操作”
func (h *MyMaxHeap) up(i int) {
	for {
		p := h.parent(i)
		if p < 0 || p > h.len() || h.data[p] > h.data[i] {
			break
		}
		if i == p {
			break
		}
		h.swap(i, p)
		i = p
	}
}

// down 从顶到底做“堆化操作”
func (h *MyMaxHeap) down(i int) {
	for {
		j := h.left(i)
		if j < 0 || j > h.len() || h.data[j] < h.data[i] {
			j = i
		}
		j1 := h.right(i)
		if j1 < 0 || j1 > h.len() || h.data[j1] > h.data[j] {
			j = j1
		}
		if j == i {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *MyMaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MyMaxHeap) len() int {
	return len(h.data)
}
