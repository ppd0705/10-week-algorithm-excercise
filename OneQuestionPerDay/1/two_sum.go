package main

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, 0)
	for i, n := range nums {
		m := target - n
		if j, ok := cache[m]; ok {
			return []int{j, i}
		}
		cache[n] = i
	}
	return nil
}
