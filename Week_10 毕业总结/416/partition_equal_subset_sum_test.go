package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target bool
	}{
		{[]int{8, 1, 2, 3}, false},
		{[]int{1, 2, 3, 6}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 3, 4}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{1, 2, 3, 15}, false},
	}

	for _, tt := range tests {
		if ans := canPartition(tt.nums); ans != tt.target {
			t.Fatalf("nums: %v, ans: %v", tt.nums, ans)
		}
	}
}
