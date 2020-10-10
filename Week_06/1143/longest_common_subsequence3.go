package main

func longestCommonSubsequence3(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		pre := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
			pre = tmp
		}
	}
	return dp[n]
}
