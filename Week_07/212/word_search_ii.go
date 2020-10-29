package main

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	trie := &Trie{children: map[byte]*Trie{}}
	m, n := len(board), len(board[0])
	for _, word := range words {
		node := trie
		for i := 0; i < len(word); i++ {
			if _, ok := node.children[word[i]]; !ok {
				node.children[word[i]] = &Trie{children: map[byte]*Trie{}}
			}
			node = node.children[word[i]]
		}
		node.isEnd = true
	}
	offset := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(i, j int, node *Trie, s string)
	dfs = func(i, j int, node *Trie, s string) {
		v := board[i][j]
		s += string(v)
		if _, ok := node.children[v]; !ok {
			return
		}
		node = node.children[v]
		if node.isEnd {
			res = append(res, s)
			node.isEnd = false
		}
		board[i][j] = '#'
		for k := 0; k < 4; k++ {
			dx, dy := offset[k][0], offset[k][1]
			x, y := i+dx, j+dy
			if x >= 0 && x < m && y >= 0 && y < n {
				dfs(x, y, node, s)
			}
		}
		board[i][j] = v

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, trie, "")
		}
	}
	return res
}
