package questionbank

import "fmt"

// BinaryFindTarget
// 给定一个长度为n的有序数组 nums ，其中可能包含重复元素。请返回数组中最左一个元素 target 的索引。若数组中不包含该元素，则返回-1.
func BinaryFindTarget(nums []int, target int) int {
	i, j := 0, len(nums)-1
	mid := (j-i)/2 + i
	for i <= j {
		fmt.Println("mid: ", mid)
		if nums[mid] == target {
			left := mid
			for left >= 0 && nums[left] == target {
				left--
			}
			return left + 1
		}
		// <---
		if nums[mid] > target {
			j = mid - 1
			mid = (j-i)/2 + i
		}
		// --->
		if nums[mid] < target {
			i = mid + 1
			mid = (j-i)/2 + i
		}
	}
	return -1
}
