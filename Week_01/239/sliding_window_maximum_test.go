package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		target []int
	}{
		{[]int{}, 3, nil},
		{[]int{1, 2, 3}, 3, []int{3}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	}

	for _, tt := range tests {
		if ans := maxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
