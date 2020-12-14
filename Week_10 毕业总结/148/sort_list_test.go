package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		vals   []int
		target []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{2, 1, 3, 7, -1}, []int{-1, 1, 2, 3, 7}},
	}
	for _, tt := range tests {
		dummy := &ListNode{}
		pre := dummy
		for _, v := range tt.vals {
			node := &ListNode{Val: v}
			pre.Next = node
			pre = node
		}
		node2 := sortList2(dummy.Next)
		ans := []int{}
		for node2 != nil {
			ans = append(ans, node2.Val)
			node2 = node2.Next
		}
		if !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("ans: %v, target: %v", ans, tt.target)
		}
	}

}
