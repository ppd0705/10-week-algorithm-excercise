package main

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	combination := make([]int, k)

	var dfs func(start, idx int)
	dfs = func(start, idx int) {
		for i := start; i <= n-k+idx+1; i++ {
			combination[idx] = i
			if idx == k-1 {
				ans = append(ans, append([]int{}, combination...))
			} else {
				dfs(i+1, idx+1)
			}
		}
	}
	dfs(1, 0)
	return ans
}
