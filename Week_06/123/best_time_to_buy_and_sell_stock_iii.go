package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	k := 2
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][k] = []int{0, 0}
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i])
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}