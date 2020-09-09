package main

func twoSum2(nums []int, target int) []int {
	cache := map[int]int{}
	for i, n := range nums {
		m := target - n
		if j, ok := cache[m]; ok {
			return []int{j, i}
		}
		cache[n] = i
	}
	return nil
}
