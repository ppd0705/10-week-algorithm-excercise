package main

func moveZeroes(nums []int) {
	count := 0
	for i, n := range  nums {
		if n == 0 {
			count++
		} else if count > 0 {
			nums[i-count], nums[i] = nums[i], 0
		}
	}
}

func moveZeroes2(nums []int) {
	j := 0
	for i, n := range nums {
		if n != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}