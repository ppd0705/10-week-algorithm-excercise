package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		heights []int
		target  int
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 0},
		{[]int{1, 0, 1}, 1},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, tt := range tests {
		if ans := trap(tt.heights); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
