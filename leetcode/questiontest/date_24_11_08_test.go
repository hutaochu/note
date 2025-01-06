package questiontest

import (
	"fmt"
	"github.com/hutaochu/note/leetcode/questionbank"
	"testing"
)

func TestMaxSumSubArray(t *testing.T) {
	nums := []int{5, 4, -1, 7, 1}
	fmt.Println("\n", questionbank.MaxSumSubArray(nums))
}
