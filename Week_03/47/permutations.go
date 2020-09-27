package main

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	elements := make(map[int]int, 0)
	for _, n := range nums {
		elements[n]++
	}

	option := make([]int, len(nums))
	var dfs func(idx int)
	dfs = func(idx int) {
		for k, v := range elements {
			if v > 0 {
				elements[k]--
				option[idx] = k
				if idx == len(nums)-1 {
					ans = append(ans, append([]int{}, option...))
				} else {
					dfs(idx + 1)
				}
				elements[k]++
			}
		}
	}
	dfs(0)
	return ans
}
