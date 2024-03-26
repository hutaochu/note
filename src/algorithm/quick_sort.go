package algorithm

func QuickSort(in []int, start, end int) {
	if start < end {
		i := parttition(in, start, end)
		QuickSort(in, start, i-1)
		QuickSort(in, i+1, end)
	}
}

func parttition(in []int, p, q int) int {
	temp := in[q]
	middleIndex := p
	for i := p; i < q; i++ {
		if in[i] <= temp {
			exchange(in, middleIndex, i)
			middleIndex++
		}
	}
	exchange(in, middleIndex, q)
	return middleIndex
}

func exchange(in []int, i, j int) {
	x := in[i]
	in[i] = in[j]
	in[j] = x
}
