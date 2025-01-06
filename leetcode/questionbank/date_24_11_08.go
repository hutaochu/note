package questionbank

// MaxSumSubArray godoc
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
// 示例 1：
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2：
//
// 输入：nums = [1]
// 输出：1
// 示例 3：
//
// 输入：nums = [5,4,-1,7,8]
// 输出：23
//
// 提示：
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
//
// 解法说明：
//
// 𝑑𝑝[𝑖]=𝑀𝑎𝑡ℎ.𝑚𝑎𝑥(𝑑𝑝[𝑖−1],0)+𝑛𝑢𝑚𝑠[𝑖]
//
// 如果以i-1为结尾的子序列最大值为负，那么不管怎么样，dp[i]就应该从头开始计数，因为前面的子序列已经不再有影响了，(dp[i - 1] + nums[i]必然小于nums[i])
func MaxSumSubArray(nums []int) int {
	maxCurrent, maxGlobal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// if maxCurrent < 0 {
		//  maxCurrent = nums[i]
		// } else {
		// 	maxCurrent = maxCurrent + nums[i]
		// }
		maxCurrent = maxInt(nums[i], maxCurrent+nums[i])
		maxGlobal = maxInt(maxCurrent, maxGlobal)
	}

	return maxGlobal
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
