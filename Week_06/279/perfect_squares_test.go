package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		n      int
		target int
	}{
		{12, 3},
	}
	for _, tt := range tests {
		if ans := numSquares(tt.n); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
