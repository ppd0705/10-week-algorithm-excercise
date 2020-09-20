package main

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	if len(arr) == k {
		return arr
	}
	quickSearch(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSearch(arr []int, left, right, k int) {
	if n := partition(arr, left, right); n > k {
		quickSearch(arr, left, n-1, k)
	} else if n < k {
		quickSearch(arr, n, right, k)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i, j := left-1, left
	for j < right {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		j++
	}
	i++
	arr[right], arr[i] = arr[i], arr[right]
	return i
}
