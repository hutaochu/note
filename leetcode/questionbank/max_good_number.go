package questionbank

import (
	"slices"
	"strconv"
)

// MaxGoodNumber
//
// 给你一个长度为 3 的整数数组 nums。
//
// 现以某种顺序 连接 数组 nums 中所有元素的 二进制表示 ，请你返回可以由这种方法形成的 最大 数值。
//
// 注意 任何数字的二进制表示 不含 前导零。
//
// 示例 1:
//
// 输入: nums = [1,2,3]
//
// 输出: 30
//
// 解释:
//
// 按照顺序 [3, 1, 2] 连接数字的二进制表示，得到结果 "11110"，这是 30 的二进制表示。
//
// 示例 2:
//
// 输入: nums = [2,8,16]
//
// 输出: 1296
//
// 解释:
//
// 按照顺序 [2, 8, 16] 连接数字的二进制表述，得到结果 "10100010000"，这是 1296 的二进制表示。
//
// 提示:
//
// nums.length == 3
// 1 <= nums[i] <= 127
//
// 注意：竞赛中，请勿复制题面内容，以免影响您的竞赛成绩真实性。
func MaxGoodNumber(nums []int) int {
	var binaryNums []string
	for _, v := range nums {
		binaryNums = append(binaryNums, strconv.FormatInt(int64(v), 2))
	}
	slices.SortFunc(binaryNums, func(a, b string) int {
		i := 0
		for ; i < len(a) && i < len(b); i++ {
			if a[i] == b[i] {
				continue
			}
		}
		if a[i-1] != b[i-1] {
			ai, _ := strconv.Atoi(string(a[i-1]))
			bi, _ := strconv.Atoi(string(b[i-1]))
			return -(ai - bi)
		}

		return -(len(b) - len(a))
	})
	var binaryNum string
	for _, v := range binaryNums {
		binaryNum += v
	}
	res, _ := strconv.ParseInt(binaryNum, 2, 0)
	return int(res)
}
