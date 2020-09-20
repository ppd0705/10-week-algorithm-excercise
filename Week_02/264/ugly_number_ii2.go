package main

func nthUglyNumber2(n int) int {
	dp := make([]int, n)
	dp[0]=1
	two := 0
	three := 0
	five := 0
	for i := 1; i < n; i++ {
		num := min(2*dp[two], 3*dp[three], 5*dp[five])
		dp[i] = num
		if 2*dp[two] == num {
			two++
		}
		if 3*dp[three] == num {
			three++
		}
		if 5*dp[five] == num {
			five++
		}
	}
	return dp[n-1]
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}
	if y <= x && y <= z {
		return y
	}
	return z
}
