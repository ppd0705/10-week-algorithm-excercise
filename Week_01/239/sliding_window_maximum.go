package main

func maxSlidingWindow(nums []int, k int) []int {
	var res, windows []int

	for i, n := range nums {
		if len(windows) > 0 && windows[0] <= i-k {
			windows = windows[1:]
		}

		for len(windows) > 0 && nums[windows[len(windows)-1]] < n {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, i)

		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}
	}
	return res
}
