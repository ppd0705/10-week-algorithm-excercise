package main

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26;i++ {
		parent[i] = i
	}

	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '=' {
			left := int(equations[i][0]-'a')
			right := int(equations[i][3]-'a')
			union(parent, left, right)
		}
	}

	for i := 0; i < len(equations); i++ {
		if equations[i][1] != '=' {
			left := int(equations[i][0]-'a')
			right := int(equations[i][3]-'a')
			if find(parent, left) == find(parent, right) {
				return false
			}
		}
	}
	return true
}

func find(parent []int, index int) int {
	if index == parent[index] {
		return index
	}
	return find(parent, parent[index])
}

func union(parent []int, left, right int) {
	parent[find(parent, left)] = find(parent, parent[right])
}

