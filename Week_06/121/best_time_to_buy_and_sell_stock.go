package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > minPrice {
			res = max(prices[i]-minPrice, res)
		} else {
			minPrice = prices[i]
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
