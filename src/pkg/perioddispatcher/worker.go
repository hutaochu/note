package perioddispatcher

func partion(arr []int, i, j int) int {
	cursor := i
	middleValue := arr[i]
	for _, x := range arr[i:j] {
		if x < middleValue {
			x, arr[cursor] = arr[cursor], x
			cursor++
		}
	}
	return cursor
}
