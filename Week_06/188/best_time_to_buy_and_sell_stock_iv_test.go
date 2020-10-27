package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		k      int
		target int
	}{
		{[]int{1,2,4,2,5,7,2,4,9,0}, 4, 15},
		{[]int{1,2}, 1, 1},
		{[]int{3, 2, 6, 5, 0, 3}, 2, 7},
	}
	for _, tt := range tests {
		if ans := maxProfit(tt.k, tt.prices); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
