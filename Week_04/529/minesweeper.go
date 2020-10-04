package main

type point struct {
	x, y int
}

func updateBoard(board [][]byte, click []int) [][]byte {
	row, column := len(board), len(board[0])
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	pointSet := map[point]bool{point{x, y}: true}
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 1, 0, 1, -1, 0, 1, -1}
	for len(pointSet) > 0 {
		newSet := make(map[point]bool, 0)
		for p := range pointSet {
			count := 0
			pointers := make([]point, 0)
			for i := 0; i < 8; i++ {
				x, y = p.x+dx[i], p.y+dy[i]
				if x >= 0 && x < row && y >= 0 && y < column {
					if board[x][y] == 'M' {
						count++
					} else if board[x][y] == 'E' {
						pointers = append(pointers, point{x, y})
					}
				}
			}
			if count == 0 {
				board[p.x][p.y] = 'B'
				for i := 0; i < len(pointers); i++ {
					newSet[pointers[i]] = true
				}
			} else {
				board[p.x][p.y] = byte('0' + count)
			}
		}
		pointSet = newSet
	}
	return board
}
