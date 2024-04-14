package main

import "fmt"

func partition(arr []int, i, j int) int {
	cursor := i
	for index := i; index < j; index++ {
		if arr[index] < arr[cursor] {
			arr[index], arr[cursor] = arr[cursor], arr[index]
			cursor++

			arr[cursor], arr[index] = arr[index], arr[cursor]
		}
	}
	return cursor
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := partition(arr, l, r)
	quickSort(arr, l, mid-1)
	quickSort(arr, mid+1, r)
}

func main() {
	arr := []int{24, 12, 22, 11, 55, 33, 11, 22, 55}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
