package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{0,1,1,1,2,2,3,3,4}, 5},
	}

	for _, tt := range tests {
		if ans := removeDuplicates(tt.nums); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
