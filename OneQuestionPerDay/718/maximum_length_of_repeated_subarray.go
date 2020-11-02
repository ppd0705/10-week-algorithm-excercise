package main

func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				res = max(res, dp[i+1][j+1])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
