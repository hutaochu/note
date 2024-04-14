package algorithm

func QuickSort(in []int, start, end int) {
	if start < end {
		i := partition(in, start, end)
		QuickSort(in, start, i-1)
		QuickSort(in, i+1, end)
	}
}

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
