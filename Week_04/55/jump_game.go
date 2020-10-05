package main

func canJump(nums []int) bool {
	n := len(nums)
	rightMost := 0

	for i := 0; i < n; i++ {
		if i <= rightMost {
			rightMost = max(rightMost, i+nums[i])
		}
		if rightMost >= n-1 {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
