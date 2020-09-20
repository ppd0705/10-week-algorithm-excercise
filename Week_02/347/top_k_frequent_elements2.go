package main

func topKFrequent2(nums []int, k int) []int {
	counter := make(map[int]int, 0)
	for _, n := range nums {
		counter[n]++
	}
	arr := make([][2]int, 0)
	for key, v := range counter {
		arr = append(arr, [2]int{key, v})
	}
	if len(arr) > k {
		quickSearch(arr, 0, len(arr)-1, k)
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][0]
	}
	return res
}

func quickSearch(arr [][2]int, left, right, k int) {
	if n := partition(arr, left, right); n > k {
		quickSearch(arr, left, n-1, k)
	} else if n < k {
		quickSearch(arr, n+1, right, k)
	}
}

func partition(arr [][2]int, left, right int) int {
	pivot := arr[right][1]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j][1] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
