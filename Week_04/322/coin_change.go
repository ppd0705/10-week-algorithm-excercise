package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		v := math.MaxInt64
		for _, c := range coins {
			if i-c >= 0 {
				v = min(v, dp[i-c]+1)
			}
		}
		dp[i] = v
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
