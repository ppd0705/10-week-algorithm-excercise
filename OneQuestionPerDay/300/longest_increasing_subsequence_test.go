package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18, 4, 8, 6, 12}, 5},
	}

	for _, tt := range tests {
		if ans := lengthOfLIS2(tt.nums); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
