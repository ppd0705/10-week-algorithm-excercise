package main

import "sort"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]) {
			return true
		}
		return false
	})

	res := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		}
	}
	return res
}
