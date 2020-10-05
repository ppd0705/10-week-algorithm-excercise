package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 2, 1, 4, 4}, 3},
		{[]int{2, 3, 1, 1, 4}, 2},
	}

	for _, tt := range tests {
		if ans := jump(tt.nums); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
