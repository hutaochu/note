package data_structure

type Queue struct {
	data     []any
	capacity int
	head     int
	tail     int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data:     make([]any, cap),
		capacity: cap,
		head:     -1,
		tail:     -1,
	}
}

func (q *Queue) Enqueue(v any) bool {
	return true
}

func (q *Queue) Dequeue() (any, bool) {
	return "", false
}
