package main

import "sort"

type Element struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	elements := make(map[Element]bool)

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		cache := map[int]int{}
		for j := i + 1; j < len(nums); j++ {
			c := nums[j]
			b := - a - c
			if _, ok := cache[b]; ok {
				elements[Element{a, b, c}] = true
			} else {
				cache[c] = 1
			}
		}
	}
	for e := range elements {
		ret = append(ret, []int{e.a, e.b, e.c})
	}
	return ret
}
