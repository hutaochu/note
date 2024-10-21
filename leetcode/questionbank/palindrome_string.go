package questionbank

// PalindromeString godoc
//
// 给你一个字符串 s，找到 s 中最长的 回文 子串。
//
// 示例 1：
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
//
// 输入：s = "cbbd"
// 输出："bb"
func PalindromeString(s string) string {
	var maxPalindromeStr string
	for i := range s {
		for j := i; j < len(s); j++ {
			s1 := getPalindromeStr(s, i, i)
			s2 := getPalindromeStr(s, i, i+1)
			if len(s1) > len(maxPalindromeStr) {
				maxPalindromeStr = s1
			}
			if len(s2) > len(maxPalindromeStr) {
				maxPalindromeStr = s2
			}
		}
	}
	return maxPalindromeStr
}

func getPalindromeStr(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// PalindromeStringDP godoc
// 设 dp[i][j] 表示字符串从索引 i 到 j 是否为回文。
//
// 若 s[i] == s[j] 且内部子串 s[i+1:j-1] 也是回文，则 dp[i][j] = true
func PalindromeStringDP(s string) string {
	return ""
}
