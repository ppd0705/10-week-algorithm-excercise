package main

import (
	"math"
)

func coinChange2(coins []int, amount int) int {
	var dfs func(n int) int
	cache := make(map[int]int, 0)
	dfs = func(n int) int {
		if n < 0 {
			return -1
		}
		if n == 0 {
			return 0
		}
		if v, ok := cache[n]; ok {
			return v
		}
		res := math.MaxInt64
		for _, c := range coins {
			if v := dfs(n - c); v >= 0 && v < res {
				res = v + 1
			}
		}
		cache[n] = res
		return res
	}
	if dfs(amount) == math.MaxInt64 {
		return -1
	}
	return dfs(amount)
}
