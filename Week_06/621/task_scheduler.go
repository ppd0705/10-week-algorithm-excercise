package main

import "sort"

func leastInterval(tasks []byte, n int) int {
	m := len(tasks)
	if m < 2 || n == 0 {
		return m
	}

	counter := make([]int, 26)
	for i := 0; i < m; i++ {
		counter[tasks[i]-'A']++
	}
	sort.Ints(counter)

	maxCount := counter[25]
	x := 0
	for i := 0; i < 25; i++ {
		if counter[i] == maxCount {
			x++
		}
	}
	return max((maxCount-1)*(n+1)+1+x, m)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
