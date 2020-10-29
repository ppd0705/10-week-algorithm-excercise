package main

func findCircleNum(M [][]int) int {
	n := len(M)
	res := 0

	visited := make([]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if !visited[j] && M[i][j] == 1 {
				visited[j] = true
				dfs(j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			res++
			visited[i] = true
			dfs(i)
		}
	}
	return res
}
