package main

func heapify(nums []int, root, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		if child < end && nums[child] < nums[child+1] {
			child++
		}
		if nums[root] >= nums[child] {
			return
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
	}
}
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n-1)
	}

	// pop k-1 times
	for i := 1; i < k; i++ {
		nums[0], nums[n-i] = nums[n-i], nums[0]
		heapify(nums, 0, n-i-1)
	}
	return nums[0]
}
