package main

import (
	"sort"
)

type NumIndex [][2]int

func (n NumIndex) Len() int {
	return len(n)
}

func (n NumIndex) Less(i, j int) bool {
	if n[i][0] <= n[j][0] {
		return true
	}
	return false
}

func (n NumIndex) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func twoSum(nums []int, target int) []int {
	var numIndex NumIndex
	for i, n := range nums {
		numIndex = append(numIndex, [2]int{n, i})
	}
	sort.Sort(numIndex)
	//fmt.Println(numIndex)
	i, j := 0, len(numIndex)-1
	for i < j {
		s := numIndex[i][0] + numIndex[j][0]
		//fmt.Println(i, j, s)
		if s == target {
			return []int{numIndex[i][1], numIndex[j][1]}
		} else if s < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
