package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	left, right := 0, m-1
	for left <= right {
		mid := (left + right) / 2
		if target < matrix[mid][0] {
			right = mid - 1
		} else if target > matrix[mid][n-1] {
			left = mid + 1
		} else {
			l, r := 0, n-1
			for l <= r {
				c := (l + r) / 2
				if target == matrix[mid][c] {
					return true
				} else if target > matrix[mid][c] {
					l = c + 1
				} else {
					r = c - 1
				}
			}
			return false
		}
	}
	return false
}
