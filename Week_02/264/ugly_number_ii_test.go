package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		n      int
		target int
	}{
		{1, 1},
		{10, 12},
	}
	for _, tt := range tests {
		if ans := nthUglyNumber2(tt.n); ans != tt.target {
			t.Fatalf("ans: %d, target: %d\n", ans, tt.target)
		}
	}
}
