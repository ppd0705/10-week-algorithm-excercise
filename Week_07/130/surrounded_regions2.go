package main

func solve2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	parent := make([]int, m*n+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	dummy := len(parent) - 1
	parent[dummy] = dummy

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && i == 0 || i == m-1 || j == 0 || j == n-1 {
				parent[i*n+j] = dummy
			} else {
				parent[i*n+j] = i*n + j
			}
		}
	}
	dp := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					continue
				} else {
					for k := range dp {
						if board[i+dp[k][0]][j+dp[k][1]] == 'O' {
							union(parent, dummy, i*n+j, (i+dp[k][0])*n+j+dp[k][1])
						}
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && find(parent, i*n+j) != dummy {
				board[i][j] = 'X'
			}
		}
	}
}

func find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int,dummy,  x, y int) {
	px := find(parent, x)
	py := find(parent, y)
	if px == dummy {
		parent[py] = px
	} else {
		parent[px] = py
	}
}
