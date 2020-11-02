package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		target int
	}{
		{[]int{1,2,3}, []int{4,2,3}, 0},
	}
	for _, tt := range tests {
		if ans := findLength2(tt.a, tt.b); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
