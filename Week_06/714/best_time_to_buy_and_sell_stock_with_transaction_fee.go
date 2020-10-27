package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int, fee int) int {
	hold, notHold := -prices[0]-fee, 0
	for i := 1; i < len(prices); i++ {
		notHold, hold = max(notHold, hold+prices[i]),max(hold, notHold-prices[i]-fee)
	}
	return notHold
}
