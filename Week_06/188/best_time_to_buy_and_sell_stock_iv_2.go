package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxProfitWithoutLimit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func maxProfit2(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k == 0 {
		return 0
	}
	if k*2 >= n {
		return maxProfitWithoutLimit(prices)
	}
	dp := make([][2]int, k+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[j][1] = -prices[i]
			} else {
				dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
				dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
			}

		}
	}

	return dp[k][0]
}
