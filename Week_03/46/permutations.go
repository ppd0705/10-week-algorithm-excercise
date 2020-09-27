package main

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	k := len(nums)
	visited := make(map[int]bool, k)
	option := make([]int, k)
	var dfs func(idx int)
	dfs = func(idx int) {
		for _, n := range nums {
			if visited[n] {
				continue
			}
			option[idx] = n
			visited[n] = true
			if idx == k-1 {
				ans = append(ans, append([]int{}, option...))
			} else {
				dfs(idx + 1)
			}
			visited[n] = false
		}
	}
	dfs(0)
	return ans
}
