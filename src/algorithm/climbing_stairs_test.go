package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	r := ClimbStairs(2)
	assert.Equal(t, 2, r, "r should be 2")
	r = ClimbStairs(3)
	assert.Equal(t, 3, r, "r should be 3")

	fmt.Println(ClimbStairsDP(100), ClimbStairsWithCache(40))
}
