package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		heights []int
		target  int
	}{
		{[]int{1, 1, 1, 1}, 3},
		{[]int{1, 1}, 1},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, tt := range tests {
		if ans := maxArea(tt.heights); ans != tt.target {
			t.Fatalf("target : %d, ans: %d\n", tt.target, ans)
		}
	}
}
