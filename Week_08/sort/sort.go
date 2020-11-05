package main

import "fmt"

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		preIndex := i - 1
		current := nums[i]
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = current
	}
}

func merge(left, right []int) []int {
	i, j := 0, 0
	res := make([]int, 0)

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for i < len(left) {
		res = append(res, left[i])
		i++
	}
	for j < len(right) {
		res = append(res, right[j])
		j++
	}
	return res
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func doPartition(nums []int) int {
	pivotIndex := 0
	pivot := nums[pivotIndex]
	j := pivotIndex + 1
	for i := j; i < len(nums); i++ {
		if nums[i] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j-1], nums[pivotIndex] = nums[pivotIndex], nums[j-1]
	return j - 1
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	partitionIndex := doPartition(nums)
	quickSort(nums[:partitionIndex])
	quickSort(nums[partitionIndex+1:])
}

func heapify(heap []int, root, end int) {
	for {
		child := root*2 + 1 //left child
		if child > end {
			return
		}
		if child+1 <= end && heap[child+1] > heap[child] {
			child = child + 1
		}
		if heap[root] > heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums)-1)
	}
	for i := len(nums)-1; i>0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, 0, i-1)
	}
}
func main() {
	nums := []int{4, 6, 5, 1, 3, 2}
	//fmt.Printf("ans: %v\n", mergeSort(nums))
	//quickSort(nums)
	heapSort(nums)
	fmt.Printf("nums: %v\n", nums)

}
