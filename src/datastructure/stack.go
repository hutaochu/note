package datastructure

type IStack interface {
	Pop() (any, bool)
	Push(v any) bool
}

type Stack struct {
	data     []any
	top      int
	capacity int
}

func NewStack(cap int) *Stack {
	s := Stack{
		data:     make([]any, cap),
		top:      -1,
		capacity: cap,
	}
	return &s
}

func (s *Stack) Pop() (any, bool) {
	// 检查是否为空
	if s.isEmpty() {
		return "", false
	}
	r := s.data[s.top]
	s.top -= 1
	return r, true
}

func (s *Stack) Push(v any) bool {
	// 检查容量是否已满
	if s.isFull() {
		return false
	}
	s.top += 1
	s.data[s.top] = v
	return true
}

func (s *Stack) isFull() bool {
	return s.top >= s.capacity
}

func (s *Stack) isEmpty() bool {
	return s.top < 0
}
