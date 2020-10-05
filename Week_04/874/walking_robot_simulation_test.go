package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		commands  []int
		obstacles [][]int
		target    int
	}{
		{[]int{4, -1, 3}, [][]int{}, 25},
		{[]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}, 65},
	}
	for _, tt := range tests {
		if ans := robotSim(tt.commands, tt.obstacles); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
