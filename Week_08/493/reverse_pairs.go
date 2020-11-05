package main

func bisectLeft(nums []int, n int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func bisectInsert(nums []int, n int) []int {
	l := bisectLeft(nums, n)
	nums = append(nums, 0)
	copy(nums[l+1:], nums[l:len(nums)-1])
	nums[l] = n
	return nums
}

func reversePairs(nums []int) int {
	res := 0
	arr := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		res += bisectLeft(arr, nums[i])
		arr = bisectInsert(arr, nums[i]*2)
	}
	return res
}
