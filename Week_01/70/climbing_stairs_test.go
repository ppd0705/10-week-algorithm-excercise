package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		n      int
		target int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, tt := range tests {
		if ans := climbStairs(tt.n); ans != tt.target {
			t.Fatalf("n: %d, ans %d, target %d\n", tt.n, ans, tt.target)
		}
	}
}
