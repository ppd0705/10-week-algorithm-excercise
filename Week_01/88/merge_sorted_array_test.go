package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	t.Helper()
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{[]int{0, 0}, 0, []int{1, 5}, 2},
		{[]int{0, 0, 0}, 1, []int{1, 5}, 2},
		{[]int{3,7, 30, 0, 0}, 3, []int{1, 5}, 2},
		{[]int{1, 4, 9, 0, 0}, 3, []int{1, 5}, 2},
	}
	for _, tt := range tests {
		target := append(tt.nums2, tt.nums1[:tt.m]...)
		sort.Ints(target)
		if merge(tt.nums1, tt.m, tt.nums2, tt.n); !reflect.DeepEqual(tt.nums1, target) {
			t.Fatalf("target: %v, ans: %v\n", target, tt.nums1)
		}

	}
}
