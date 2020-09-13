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
		{[]int{1}, 0, []int{1}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
	}

	for _, tt := range tests {
		if rotate(tt.nums, tt.k); !reflect.DeepEqual(tt.nums, tt.target) {
			t.Fatalf("target %v, got %v\n", tt.target, tt.nums)
		}
	}
}
