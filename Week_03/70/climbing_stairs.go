package main

var cache = map[int]int{}

func climbStairs(n int) int {
	if v, ok := cache[n]; ok {
		return v
	} else {
		if n < 3 {
			v = n
		} else {
			v = climbStairs(n-2) + climbStairs(n-1)
		}
		cache[n] = v
		return v
	}
}
