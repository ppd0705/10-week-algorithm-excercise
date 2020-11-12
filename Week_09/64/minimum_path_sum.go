package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])
	if n == 0 {
		return -1
	}
	dp := make([][]int, m)
	for i := 0;i < m;i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j:= 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}