package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		n      int
		target int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
	}

	for _, tt := range tests {
		if ans := mySqrt(tt.n); ans != tt.target {
			t.Fatalf("n: %d, target: %d, ans: %d\n", tt.n, tt.target, ans)
		}
	}
}
