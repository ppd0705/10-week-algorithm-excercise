package main

type point struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	if grid[0][0] == 1 || grid[row-1][column-1] == 1 {
		return -1
	}

	queue := []point{point{0, 0}}
	res := 0
	offsets := [8]point{
		{0, 1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{1, 0},
		{1, -1},
		{1, 1},
	}
	for len(queue) > 0 {
		res++
		l := len(queue)
		for i := 0; i < l; i++ {
			x, y := queue[i].x, queue[i].y
			if x == row-1 && y == column-1 {
				return res
			}
			for _, o := range offsets {
				nextX, nextY := x+o.x, y+o.y
				if nextX >= 0 && nextX < row && nextY >= 0 && nextY < column && grid[nextX][nextY] == 0 {
					queue = append(queue, point{nextX, nextY})
					grid[nextX][nextY] = 1
				}
			}
		}
		queue = queue[l:]
	}
	return -1
}
