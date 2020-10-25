package main

func maxProduct(nums []int) int {
	ans, maxF, minF := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		m, n := maxF, minF
		maxF = max(m*nums[i], max(n*nums[i], nums[i]))
		minF = min(m*nums[i], min(n*nums[i], nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
