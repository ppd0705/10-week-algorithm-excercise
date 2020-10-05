package main

type point struct {
	x, y int
}

func robotSim(commands []int, obstacles [][]int) int {
	dx, dy := 0, 1
	curr := point{0, 0}
	obstaclesMap := make(map[point]bool, 0)
	for i := 0; i < len(obstacles); i++ {
		obstaclesMap[point{obstacles[i][0], obstacles[i][1]}] = true
	}
	res := 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == -2 {
			dx, dy = -dy, dx
		} else if commands[i] == -1 {
			dx, dy = dy, -dx
		} else {
			x, y := curr.x, curr.y
			for j := 0; j < commands[i]; j++ {
				x += dx
				y += dy
				if obstaclesMap[point{x, y}] {
					x -= dx
					y -= dy
					break
				}
			}
			curr.x, curr.y = x, y
			distance := curr.x*curr.x + curr.y*curr.y
			if distance > res {
				res = distance
			}
		}
	}
	return res
}
