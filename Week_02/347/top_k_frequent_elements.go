package main

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, 0)
	for _, n := range nums {
		counter[n]++
	}
	data := make([][2]int, k)

	i := 0
	for key, v := range counter {
		if i <= k-1 {
			data[i] = [2]int{key, v}
			if i == k-1 {
				for j := k/2 - 1; j >= 0; j-- {
					heaplify(data, j, k-1)
				}
			}
			i++
		} else if v > data[0][1] {
			data[0][0] = key
			data[0][1] = v
			heaplify(data, 0, k-1)
		}
	}
	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = data[i][0]
	}
	return res
}

func heaplify(arr [][2]int, root int, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		if child+1 <= end && arr[child][1] > arr[child+1][1] {
			child++
		}
		if arr[root][1] < arr[child][1] {
			return
		}
		arr[root], arr[child] = arr[child], arr[root]
		root = child
	}
}
