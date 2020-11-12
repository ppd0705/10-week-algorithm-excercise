package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	left := make([]int, n)
	for j := 0; j < n; j++ {
		left[j] = -1
	}
	right := make([]int, n)
	for j := 0; j < n; j++ {
		right[j] = n
	}
	height := make([]int, n)

	res := 0

	for i := 0; i < m; i++ {
		curLeft, curRight := -1, n
		//height
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		//left
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				left[j] = max(left[j], curLeft)
			} else {
				left[j] = -1
				curLeft = j
			}
		}
		//right
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = min(right[j], curRight)
			} else {
				right[j] = n
				curRight = j
			}
		}
		//area
		for j := 0; j < n; j++ {
			res = max(res, height[j]*(right[j]-left[j]-1))
		}
		//fmt.Printf("%d:  \n%v\n%v\n", i, left, right)
		//for j := 0; j < n; j++ {
		//	fmt.Printf("j: %d, left: %d, right: %d, area: %d\n", j, left[j], right[j], height[j]*(right[j]-left[j]-1))
		//}
	}
	return res
}
