package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target []int
	}{
		{[]int{1, 2, 0}, []int{1, 2, 0}},
		{[]int{1, 0, 0}, []int{1, 0, 0}},
		{[]int{1, 0, 0, 2, 0}, []int{1, 2, 0, 0, 0}},
	}

	for _, tt := range tests {
		if moveZeroes(tt.nums); !reflect.DeepEqual(tt.nums, tt.target) {
			t.Fatalf("target: %v, nums: %v\n", tt.target, tt.nums)
		}
	}
}
