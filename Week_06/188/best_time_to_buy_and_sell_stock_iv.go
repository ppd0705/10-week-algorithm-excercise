package main

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k == 0 {
		return 0
	}
	if k * 2 >= n {
		return maxProfitWithoutLimit(prices)
	}
	dp1 := make([][2]int, k+1)
	dp2 := make([][2]int, k+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp2[j][1] = -prices[i]
			} else {
				dp2[j][0] = max(dp1[j][0], dp1[j][1]+prices[i])
				dp2[j][1] = max(dp1[j][1], dp1[j-1][0]-prices[i])
			}

		}
		copy(dp1, dp2)
	}

	return dp2[k][0]
}