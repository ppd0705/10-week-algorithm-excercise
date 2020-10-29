package main

func isValidSudoku(board [][]byte) bool {
	var row, column, box [9]uint16
	var cur uint16
	for i := 0; i < 9;i++ {
		for j := 0;j < 9;j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')
			k := i/3*3 + j/3
			if (row[i]&cur) | (column[j]&cur) | (box[k]&cur) != 0{
				return false
			}
			row[i] |= cur
			column[j] |= cur
			box[k] |= cur
		}
	}
	return true
}