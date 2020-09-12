package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		heights []int
		target  int
	}{
		{[]int{}, 0},
		{[]int{2, 1, 5, 6, 2, 3}, 10},
	}

	for _, tt := range tests {
		if ans := largestRectangleArea(tt.heights); ans != tt.target {
			t.Fatalf("target: %d, ans： %d\n", tt.target, ans)
		}
	}
}
