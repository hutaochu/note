package datastructure

type CycleArray[T any] struct {
	list     []T
	head     int
	tail     int
	capacity int
	size     int
}

func NewCycleArray[T any](capacity int) CycleArray[T] {
	return CycleArray[T]{
		list:     make([]T, capacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: capacity,
	}
}

func (c *CycleArray[T]) AddFirst(e T) {
	if c.isFull() {
		c.resize()
	}

	c.head = (c.head - 1 + c.capacity) % c.capacity
	c.list[c.head] = e
	c.size += 1
}

func (c *CycleArray[T]) resize() {}

func (c *CycleArray[T]) isFull() bool {
	return c.capacity == c.size
}

func (c *CycleArray[T]) isEmpty() bool {
	return c.size == 0
}
