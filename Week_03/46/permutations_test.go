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
		{[]int{2}, [][]int{{2}}},
		{[]int{2, 3, 4}, [][]int{{2, 3, 4}, {2, 4, 3}, {3, 2, 4}, {3, 4, 2}, {4, 2, 3}, {4, 3, 2}}},
	}
	for _, tt := range tests {
		if ans := permute(tt.nums); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
