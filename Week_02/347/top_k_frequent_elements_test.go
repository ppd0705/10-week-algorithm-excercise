package main

import (
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
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		ans := topKFrequent2(tt.nums, tt.k)
		sort.Ints(ans)
		if !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans : %v\n", tt.target, ans)
		}
	}
}
