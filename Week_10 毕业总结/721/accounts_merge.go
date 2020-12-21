package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	parent := make([]int, len(accounts))
	res := make([][]string, 0)
	indexMap := make(map[string]int, 0)
	for i := 0; i < len(accounts); i++ {
		parent[i] = i
	}
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			if k, ok := indexMap[account[j]]; ok {
				union(parent, k, i)
			} else {
				indexMap[account[j]] = i
			}
		}
	}
	resIndex := make(map[int]int, 0)
	for i := 0; i < len(accounts); i++ {
		p := find(parent, i)
		if idx, ok := resIndex[p]; ok {
			res[idx] = append(res[idx], accounts[i][1:]...)
		} else {
			resIndex[p] = len(res)
			res = append(res, accounts[i])
		}
	}

	for i := 0; i < len(res); i++ {
		sort.Strings(res[i][1:])
		for j := 1; j < len(res[i])-1; {
			if res[i][j] == res[i][j+1] {
				copy(res[i][j:len(res[i])-1], res[i][j+1:])
				res[i] = res[i][:len(res[i])-1]
			} else {
				j++
			}

		}

	}
	return res
}

func find(parent []int, x int) int {
	for x != parent[x] {
		x, parent[x] = parent[x], parent[parent[x]]
	}
	return x
}

func union(parent []int, x, y int) {
	parent[find(parent, x)] = find(parent, y)
}
