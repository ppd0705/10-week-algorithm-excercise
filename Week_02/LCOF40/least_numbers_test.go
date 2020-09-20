package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		target []int
	}{
		{[]int{0,0,2,3,2,1,1,2,0,4}, 10, []int{0,0,2,3,2,1,1,2,0,4}},
		{[]int{0,0,1,2,4,2,2,3,1,4}, 8, []int{0, 0, 1, 1, 2, 2, 2, 3}},
		{[]int{1, 2, 4, 3}, 3, []int{1, 2, 3}},
		{[]int{4, 1, 3, 2}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		fmt.Printf("nums: %v\n", tt.nums)
		ans := getLeastNumbers2(tt.nums, tt.k)
		sort.Ints(ans)
		sort.Ints(tt.target)
		if !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target %v, ans: %v\n", tt.target, ans)
		}
	}
}
