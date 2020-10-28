package main

func canCross2(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			k := stones[i]-stones[j]
			if k < 1 || k > j+1 || !dp[j][k]{
				continue
			}
			dp[i][k] = dp[j][k] || dp[j][k-1] || dp[j][k+1]
			if i == n-1  && dp[i][k] {
				return true
			}
		}
	}
	return false
}
