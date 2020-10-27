package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		fee    int
		prices []int
		target int
	}{
		{1, []int{1, 3, 5, 7}, 5},
	}
	for _, tt := range tests {
		if ans := maxProfit(tt.prices, tt.fee); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
