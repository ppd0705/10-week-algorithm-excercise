package main

func majorityElement(nums []int) int {
	counter := make(map[int]int, 0)
	for _, n := range nums {
		counter[n]++
	}
	majorKey, majorValue := 0, 0
	for k, v := range counter {
		if v > majorValue {
			majorKey, majorValue = k, v
		}
	}
	return majorKey
}
