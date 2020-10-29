package main

func boxIdx(i, j int) int {
	return i/3*3 + j/3
}

func solveSudoku(board [][]byte) {
	row := [9][9]bool{}
	column := [9][9]bool{}
	box := [9][9]bool{}
	needFill := [9][9]bool{}

	needFillCount := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				needFillCount++
				needFill[i][j] = true
			} else {
				v := board[i][j] - '1'
				row[i][v] = true
				column[j][v] = true
				box[boxIdx(i, j)][v] = true
			}
		}
	}
	getNext := func(i, j int) (int, int) {
		if j == 8 {
			return i+1, 0
		}
		return i,j+1
	}
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if needFillCount == 0 {
			return true
		}
		if !needFill[x][y] {
			return dfs(getNext(x, y))
		} else {
			for v := 0; v < 9;v++ {
				if !row[x][v] && !column[y][v] && !box[boxIdx(x,y)][v] {
					row[x][v] = true
					column[y][v] = true
					box[boxIdx(x,y)][v] = true
					board[x][y] = byte(v + '1')
					needFillCount--
					if dfs(getNext(x,y)) {
						return true
					}
					row[x][v] = false
					column[y][v] = false
					box[boxIdx(x,y)][v] = false
					board[x][y] = '.'
					needFillCount++
				}
			}
		}
		return false
	}
	dfs(0, 0)
}
