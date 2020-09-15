package main

func intersect(nums1 []int, nums2 []int) []int {
	counter1 := make(map[int]int, 0)
	counter2 := make(map[int]int, 0)

	for _, n := range nums1 {
		counter1[n]++
	}
	for _, n := range nums2 {
		counter2[n]++
	}
	res := make([]int, 0)
	for k, v1 := range counter1 {
		if v2, ok := counter2[k]; ok {
			for i := 0; i < min(v1, v2); i++ {
				res = append(res, k)
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
