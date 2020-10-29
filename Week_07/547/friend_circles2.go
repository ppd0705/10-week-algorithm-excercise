package main

func findCircleNum2(M [][]int) int {
	n := len(M)
	res := 0

	parent := make([]int, n)
	for i := 0; i < n;i++ {
		parent[i] = -1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && M[i][j] == 1 {
				union(parent, i, j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if parent[i] == -1 {
			res++
		}
	}
	return res
}

func union(parent []int, x, y int) {
	px := find(parent, x)
	py := find(parent, y)
	if px != py {
		parent[px] = py
	}
}
func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}
