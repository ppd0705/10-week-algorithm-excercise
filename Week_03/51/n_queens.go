package main

func solveNQueens(n int) [][]string {
	rowPos := make([]int, n)
	columnUsed := make([]bool, n)
	leftDiagonalUsed := make(map[int]bool, n*2)  // row + col
	rightDiagonalUsed := make(map[int]bool, n*2) // row - col
	res := make([][]string, 0)

	printQueen := func() []string {
		template := make([]byte, n)
		queen := make([]string, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if rowPos[i] == j {
					template[j] = 'Q'
				} else {
					template[j] = '.'
				}
			}
			queen[i] = string(template)
		}
		return queen
	}
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res = append(res, printQueen())
			return
		}

		for column := 0; column < n; column++ {
			if columnUsed[column] || leftDiagonalUsed[row+column] || rightDiagonalUsed[row-column] {
				continue
			}
			rowPos[row] = column
			columnUsed[column] = true
			leftDiagonalUsed[row+column] = true
			rightDiagonalUsed[row-column] = true
			dfs(row + 1)
			columnUsed[column] = false
			leftDiagonalUsed[row+column] = false
			rightDiagonalUsed[row-column] = false
		}
	}
	dfs(0)
	return res
}
