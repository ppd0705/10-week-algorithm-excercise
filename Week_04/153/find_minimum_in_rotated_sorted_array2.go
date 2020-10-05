package main

func findMin2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	if nums[left] <= nums[right] {
		// no rotated
		return nums[left]
	}

	for left <= right {
		mid := (left+right)/2
		if nums[mid] > nums[mid+1] {
			return mid+1
		}
		if nums[mid] < nums[mid-1] {
			return mid
		}
		if nums[mid] > nums[0] {
			 left = mid+1
		} else {
			right =mid-1
		}
	}
	return -1
}
