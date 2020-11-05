package main

func lengthOfLIS2(nums []int) int {
	tail := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
		} else {
			left, right := 0, len(tail)
			for left < right {
				mid := (left + right) / 2
				if tail[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			tail[left] = nums[i]
		}
	}
	return len(tail)
}
