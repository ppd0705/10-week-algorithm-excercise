package main

func updateBoard2(board [][]byte, click []int) [][]byte {
	row, column := len(board), len(board[0])
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 1, 0, 1, -1, 0, 1, -1}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		count := 0
		for i := 0; i < 8; i++ {
			x, y := r+dx[i], c+dy[i]
			if x >= 0 && x < row && y >= 0 && y < column && board[x][y] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[r][c] = byte('0' + count)
		} else {
			board[r][c] = 'B'
			for i := 0; i < 8; i++ {
				x, y := r+dx[i], c+dy[i]
				if x >= 0 && x < row && y >= 0 && y < column && board[x][y] == 'E' {
					dfs(x, y)
				}
			}
		}
	}
	dfs(x, y)
	return board
}
