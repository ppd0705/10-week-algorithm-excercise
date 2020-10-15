package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int{
	pre, cur := 0, 0
	for _, n := range nums {
		pre, cur = cur, max(cur, pre+n)
	}
	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}