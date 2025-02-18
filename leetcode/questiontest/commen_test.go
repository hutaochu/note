package questiontest

import (
	"fmt"
	"testing"

	"github.com/hutaochu/note/leetcode/questionbank"
)

func TestSumTarget(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 6, 8, 9}
	res := questionbank.SumTarget(nums, 15)
	fmt.Println(res)
}

func TestLongestPalindromeSubStr(t *testing.T) {
	res := questionbank.LongestPalindromeSubStr("abcaacdabcdcbaaaaaaaaaaa")
	fmt.Println(res)
}

func TestFindMinCoveringSubStr(t *testing.T) {
	s := "ADOBECODEBANC"
	s1 := "ABC"
	fmt.Println(questionbank.FindMinCoveringSubStr(s, s1))
	s = "a"
	s1 = "a"
	fmt.Println(questionbank.FindMinCoveringSubStr(s, s1))
	s1 = "aa"
	fmt.Println(questionbank.FindMinCoveringSubStr(s, s1))
}

func TestFindSubStr(t *testing.T) {
	s := "ADOBECODEBANC"
	s1 := "NC"
	fmt.Println(questionbank.FindSubStr(s1, s))
}

func TestCoinsChange(t *testing.T) {
	fmt.Println(questionbank.CoinsChange([]int{1, 2, 5}, 11))
	fmt.Println(questionbank.CoinsChange([]int{1, 2147483647}, 2))
}
