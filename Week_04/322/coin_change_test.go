package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		target int
	}{
		{[]int{}, 0, 0},
		{[]int{3}, 2, -1},
		{[]int{3, 1}, 2, 2},
		{[]int{10, 9, 1}, 18, 2},
	}
	for _, tt := range tests {
		if ans := coinChange2(tt.coins, tt.amount); ans != tt.target {
			t.Fatalf("coins: %v, tatget: %d, ans: %d\n", tt.coins, tt.target, ans)
		}
	}
}
