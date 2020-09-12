package main

import "testing"

func Test(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, n1}
	n3 := &ListNode{3, n2}
	n4 := &ListNode{4, n3}
	n5 := &ListNode{5, n4}
	n6 := &ListNode{6, n5}

	if detectCycle(n1) != nil {
		t.Fatalf("failed")
	}

	if detectCycle(n6) != nil {
		t.Fatalf("failed")
	}
	n1.Next = n4
	if detectCycle(n5) != n4 {
		t.Fatalf("failed")
	}
}
