package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target [][]int
	}{
		{[]int{0,0,0}, [][]int{{0,0,0}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, tt := range tests {
		if ans := threeSum(tt.nums); !reflect.DeepEqual(tt.target, ans) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}

}
