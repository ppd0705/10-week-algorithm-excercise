package main

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	squares := make([]int, 0)
	for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
		squares = append(squares, i*i)
	}
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
		for _, j := range squares {
			if i >= j {
				dp[i] = min(dp[i], dp[i-j]+1)
			}
		}
	}
	return dp[n]
}
