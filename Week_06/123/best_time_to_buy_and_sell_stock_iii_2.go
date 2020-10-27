package main

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// first day status: first buy, first sold, second buy, second sold
	dp11 := -prices[0]
	dp10 := 0
	dp21 := -prices[0]
	dp20 := 0

	for i := 1; i < len(prices); i++ {
		dp11 = max(dp11, -prices[i])
		dp10 = max(dp10, dp11+prices[i])
		dp21 = max(dp21, dp10-prices[i])
		dp20 = max(dp20, dp21+prices[i])
	}
	return dp20
}
