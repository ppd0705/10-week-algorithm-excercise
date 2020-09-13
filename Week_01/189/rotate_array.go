package main

func rotate1(nums []int, k int) {
	if len(nums) < 2 || len(nums) == k {
		return
	}
	for i := 0; i < k; i++ {
		tail := nums[len(nums)-1]
		copy(nums[1:], nums[0:len(nums)-1])
		nums[0] = tail
	}
}
