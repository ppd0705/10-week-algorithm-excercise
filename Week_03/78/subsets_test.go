package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{}, {1}}},
		{[]int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
	}
	for _, tt := range tests {
		if ans := subsets(tt.nums); len(ans) != len(tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
