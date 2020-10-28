package main

func countSubstrings(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[i-1][j+1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}
