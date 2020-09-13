package main

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	preVal := nums[0]
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != preVal {
			nums[idx] = nums[i]
			idx++
			preVal = nums[i]
		}
	}
	return idx
}
