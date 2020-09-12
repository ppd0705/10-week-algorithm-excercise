package main

import "testing"

func Test(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2, Next: n1}
	n3:= &ListNode{Val: 3, Next: n2}
	n4 := &ListNode{Val: 4, Next: n3}

	if hasCycle(n1) {
		t.Fatalf("failed")
	}
	if hasCycle(n2) {
		t.Fatalf("failed")
	}
	if hasCycle(n3) {
		t.Fatalf("failed")
	}
	if hasCycle(n4) {
		t.Fatalf("failed")
	}

	n1.Next = n2
	if !hasCycle(n4) {
		t.Fatalf("failed")
	}
}

func Test2(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2, Next: n1}
	n3:= &ListNode{Val: 3, Next: n2}
	n4 := &ListNode{Val: 4, Next: n3}
	n1.Next = n3
	if !hasCycle(n4) {
		t.Fatalf("failed")
	}
}