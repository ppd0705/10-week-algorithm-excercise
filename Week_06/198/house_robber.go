package main

func rob(nums []int) int {
	pre, cur := 0, 0
	for _, n := range nums {
		cur, pre = max(cur, pre+n), cur
	}
	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
