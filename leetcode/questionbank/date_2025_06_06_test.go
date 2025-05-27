package questionbank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}

func TestMergeV2(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	mergeV2(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}
