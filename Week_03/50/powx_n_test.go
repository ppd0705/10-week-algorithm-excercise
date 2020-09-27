package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		x      float64
		n      int
		target float64
	}{
		{2, 10, 1024},
		{-2, -2, 0.25},
	}

	for _, tt := range tests {
		if ans := myPow(tt.x, tt.n); ans != tt.target {
			t.Fatalf("target: %f, ans: %f\n", tt.target, ans)
		}
	}
}
