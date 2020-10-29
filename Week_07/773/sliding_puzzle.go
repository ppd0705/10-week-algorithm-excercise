package main

func slidingPuzzle(board [][]int) int {
	target := "123450"
	neighbors := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1, 5},
		3: {0, 4},
		4: {1, 3, 5},
		5: {2, 4},
	}
	m, n := len(board), len(board[0])
	template := make([]byte, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			template[i*n+j] = byte('0'+ board[i][j])
		}
	}
	curr := string(template)
	queue := []string{curr}
	visited := map[string]bool{curr: true}
	res := -1

	genNewState := func(s string, old, new int) string {
		bytes := []byte(s)
		bytes[old], bytes[new] = bytes[new], bytes[old]
		return string(bytes)
	}
	for len(queue) > 0 {
		res++
		l := len(queue)
		for i := 0; i < l; i++ {
			s := queue[i]
			if s == target {
				return res
			}
			zeroIndex := 0
			for j := 1; j < len(s); j++ {
				if s[j] == '0' {
					zeroIndex = j
					break
				}
			}
			for _, n := range neighbors[zeroIndex] {
				newState := genNewState(s, zeroIndex, n)
				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, newState)
				}
			}
		}
		queue = queue[l:]
	}
	return -1

}
