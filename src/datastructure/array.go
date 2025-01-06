package datastructure

import "errors"

const (
	defaultSize = 1
	minSize     = 1024
)

var (
	ErrOutIndex = errors.New("out of index, please check your index")
)

type Array[T any] struct {
	list []T
	// current number of elements
	size int
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{
		list: make([]T, defaultSize),
		size: 0,
	}
}

func NewArrayWithCapacity[T any](cap int) *Array[T] {
	return &Array[T]{
		list: make([]T, cap),
		size: 0,
	}
}

func (a *Array[T]) Get(index int) (T, error) {
	var res T
	if a.isOutOfIndex(index) {
		return res, ErrOutIndex
	}
	return a.list[index], nil
}

func (a *Array[T]) Length() int {
	return a.size
}

func (a *Array[T]) Append(e T) {
	index := a.size
	if a.isOutOfIndex(index) {
		// scale up
		a.scaleUp()
		a.list[a.size] = e
		return
	}
	a.list[a.size] = e
	a.size++
}

func (a *Array[T]) scaleUp() {
	if a.size < minSize {
		newList := make([]T, 2*a.size)
		a.copyToNewList(newList)
		a.list = newList
		return
	}
	newList := make([]T, a.size+a.size/4)
	a.copyToNewList(newList)
	a.list = newList
}

func (a *Array[T]) copyToNewList(newList []T) {
	for i, v := range a.list {
		newList[i] = v
	}
}

func (a *Array[T]) isOutOfIndex(index int) bool {
	return index < 0 || index >= len(a.list)
}
