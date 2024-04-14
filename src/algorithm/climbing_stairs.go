package algorithm

var resCache map[int]int

func init() {
	resCache = make(map[int]int)
}

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

func ClimbStairsWithCache(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if r, ok := resCache[n]; ok {
		return r
	}
	return ClimbStairsWithCache(n-1) + ClimbStairsWithCache(n-2)
}

func ClimbStairsDP(n int) int {
	an := make([]int, n)
	an[0] = 1
	an[2] = 2
	for i := 3; i <= n; i++ {
		an[i] = an[i-1] + an[i-2]
	}
	return an[n]
}

func ClimbStairsDP2(n int) int {
	a1, a2 := 1, 2
	for i := 3; i <= n; i++ {
		a1, a2 = a2, a1+a2
	}
	return a2
}
