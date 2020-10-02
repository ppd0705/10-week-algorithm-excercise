package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res := 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = '0'
		if x-1 >= 0 && grid[x-1][y] == '1' {
			dfs(x-1, y)
		}
		if x+1 < m && grid[x+1][y] == '1' {
			dfs(x+1, y)
		}
		if y-1 >= 0 && grid[x][y-1] == '1' {
			dfs(x, y-1)
		}
		if y+1 < n && grid[x][y+1] == '1' {
			dfs(x, y+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
