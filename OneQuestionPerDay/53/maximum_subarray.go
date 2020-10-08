package main

func maxSubArray(nums []int) int {
	cur := nums[0]
	res := cur
	for i := 1; i < len(nums); i++ {
		cur = max(cur, 0) + nums[i]
		res = max(res, cur)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
