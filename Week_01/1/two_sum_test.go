package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		sum    int
		target []int
	}{
		{[]int{1, 3, 5}, 4, []int{0, 1}},
		{[]int{1, 3, 5, 2}, 3, []int{0, 3}},
	}

	for _, tt := range tests {
		if ans := twoSum(tt.nums, tt.sum); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans %v\n", tt.target, ans)
		}
	}
}
