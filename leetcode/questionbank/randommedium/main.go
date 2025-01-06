package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("question number: %d", rand.Intn(3368))
}

func maxSum(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxCurrent, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxCurrent < 0 {
			maxCurrent = nums[i]
		} else {
			maxCurrent = maxCurrent + nums[i]
		}

		maxSum = maxInt(maxCurrent, maxSum)
	}
	return maxSum
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
