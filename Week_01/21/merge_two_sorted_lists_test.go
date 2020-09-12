package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	l4 := &ListNode{Val: 4, Next: nil}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	r5 := &ListNode{Val: 5}
	r3 := &ListNode{Val: 3, Next: r5}
	r1 := &ListNode{Val: 1, Next: r3}
	head := mergeTwoLists2(l1, r1)

	for head != nil {
		fmt.Printf("%d --> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
