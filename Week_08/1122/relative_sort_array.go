package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counter := make([]int, 1001)
	for _, i := range arr2 {
		counter[i]++
	}
	idx := 0
	for _, i := range arr2 {
		for counter[i] > 0 {
			arr1[idx] = i
			idx++
			counter[i]--
		}
	}
	for i := range counter {
		for counter[i] > 0 {
			arr1[idx] = i
			idx++
			counter[i]--
		}
	}
	return arr1
}
