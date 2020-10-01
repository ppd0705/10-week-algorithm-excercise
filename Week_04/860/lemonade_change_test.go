package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		bills  []int
		target bool
	}{
		{[]int{10}, false},
		{[]int{5, 10, 10}, false},
		{[]int{5, 5, 5, 5, 10, 20, 10}, true},
	}

	for _, tt := range tests {
		if ans := lemonadeChange(tt.bills); ans != tt.target {
			t.Fatalf("bills: %v, target: %T, ans: %T\n", tt.bills, tt.target, ans)
		}
	}
}
