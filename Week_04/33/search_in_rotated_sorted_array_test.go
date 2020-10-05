package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		n      int
		target int
	}{
		{[]int{1}, 0, -1},
		{[]int{}, 0, -1},
		{[]int{1,2,3, 0}, 0, 3},
		{[]int{1,2,3, 0}, 3, 2},
		{[]int{1,2,3, 0}, 2, 1},
		{[]int{1,2,3, 0}, 1, 0},
	}
	for _, tt := range tests {
		if ans := search(tt.nums, tt.n); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
