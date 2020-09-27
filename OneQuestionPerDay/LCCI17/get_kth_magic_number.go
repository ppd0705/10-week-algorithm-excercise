package main

func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	three := 0
	five := 0
	seven := 0

	for i := 1; i < k; i++ {
		dp[i] = min3(3*dp[three], 5*dp[five], 7*dp[seven])
		if 3*dp[three] == dp[i] {
			three++
		}
		if 5*dp[five] == dp[i] {
			five++
		}
		if 7*dp[seven] == dp[i] {
			seven++
		}
	}
	return dp[k-1]
}

func min3(x, y, z int) int {
	return min(x, min(y, z))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
