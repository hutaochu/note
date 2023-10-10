package datastructure

import (
	"reflect"
)

type Array struct {
	list []any
}

func NewArray(length int) Array {
	arr := Array{
		list: make([]any, length),
	}
	return arr
}

func (a *Array) Capacity() int {
	return cap(a.list)
}

func (a *Array) Length() int {
	return len(a.list)
}

func (a *Array) Set(value any, index int) {
	a.list[index] = value
}

func (a *Array) Get(index int) any {
	return a.list[index]
}

func (a *Array) Insert(index int, value any) {
	if index < 0 {
		return
	}
	for i := len(a.list) - 1; i > index; i-- {
		a.list[i] = a.list[i-1]
	}
	a.list[index] = value
}

func (a *Array) Remove(index int) {
	length := len(a.list)
	if index < 0 || index >= length {
		return
	}
	for i := index; i < length-1; i++ {
		a.list[i] = a.list[i+1]
	}
	a.list[length-1] = nil
}

func (a *Array) FindOne(value any) int {
	for i := 0; i < len(a.list); i++ {
		if reflect.DeepEqual(a.list[i], value) {
			return i
		}
	}
	return -1
}

func (a *Array) FindAll(value any) []int {
	var targets []int
	for i := 0; i < len(a.list); i++ {
		if reflect.DeepEqual(a.list[i], value) {
			targets = append(targets, i)
		}
	}
	return targets
}

func (a *Array) Each() []any {
	res := make([]any, len(a.list))
	copy(res, a.list)
	return res
}

func (a *Array) Extend(enlarge int) {
	target := len(a.list) + enlarge
	res := make([]any, target)
	copy(res, a.list)
	a.list = res
}
