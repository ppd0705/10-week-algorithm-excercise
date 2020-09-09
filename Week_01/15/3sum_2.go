package main

import "sort"

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			s := nums[i] + b + c
			if s == 0 {
				ret = append(ret, []int{nums[i], b, c})
				l++
				r--
				for r > 0 && nums[r] == c {
					r--
				}
				for l < len(nums) && nums[l] == b {
					l++
				}
			} else if s > 0 {
				r--
				for r > 0 && nums[r] == c {
					r--
				}
			} else {
				l++
				for l < len(nums) && nums[l] == b {
					l++
				}
			}
		}
	}

	return ret
}
