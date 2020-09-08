package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		digits []int
		target []int
	}{
		{[]int{0}, []int{1}},
		{[]int{1, 0}, []int{1, 1}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for _, tt := range tests {
		if ans := plusOne(tt.digits); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target : %v, ans: %v\n", tt.target, ans)
		}
	}
}
