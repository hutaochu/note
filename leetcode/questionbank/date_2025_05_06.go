package questionbank

import "container/heap"

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//
// 示例 1：
//
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
// 示例 2：
//
// 输入：nums1 = [1], m = 1, nums2 = [], n = 0
// 输出：[1]
// 解释：需要合并 [1] 和 [] 。
// 合并结果是 [1] 。
// 示例 3：
//
// 输入：nums1 = [0], m = 0, nums2 = [1], n = 1
// 输出：[1]
// 解释：需要合并的数组是 [] 和 [1] 。
// 合并结果是 [1] 。
// 注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 解法1: 小顶堆
	// 解法2: 双指针
	p1, p2 := 0, m-1
	for i := 0; i < n; {
		if nums1[p1] < nums2[i] && nums1[p1] > 0 {
			p1++
			continue
		} else {
			tmp := p2
			p2++
			for tmp >= p1 {
				nums1[tmp+1] = nums1[tmp]
				tmp--
			}
			nums1[p1] = nums2[i]
			i++
			p1++
		}
	}
}

func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	h := &intMinHeap{
		data: make([]int, 0),
	}
	heap.Init(h)
	for i := 0; i < m; i++ {
		heap.Push(h, nums1[i])
	}

	for i := 0; i < n; i++ {
		heap.Push(h, nums2[i])
	}

	for j := 0; j < m+n; j++ {
		e := heap.Pop(h)
		nums1[j] = e.(int)
	}
}

type intMinHeap struct {
	data []int
}

func (h intMinHeap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h intMinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h intMinHeap) Len() int {
	return len(h.data)
}

func (h *intMinHeap) Pop() any {
	e := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return e
}

func (h *intMinHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}
