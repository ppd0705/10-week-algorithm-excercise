package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T)  {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, n1}
	n3 := &ListNode{3, n2}
	n4 := &ListNode{4, n3}
	n5 := &ListNode{5, n4}

	k := 6
	head := reverseKGroup(n5, k)
	for head != nil {
		fmt.Printf("%d-->", head.Val)
		head = head.Next
	}
}