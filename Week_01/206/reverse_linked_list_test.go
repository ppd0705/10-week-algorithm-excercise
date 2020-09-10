package main

import "testing"

func Test(t *testing.T)  {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2, Next: &node1}
	node3 := ListNode{Val: 3, Next: &node2}
	node4 := ListNode{Val: 3, Next: &node3}
	node5 := ListNode{Val: 5, Next: &node4}
	if ans := reverseList(&node1); ans != &node1 {
		t.Fatalf("failed, got %v\n", (*ans).Val)
	}

	if ans := reverseList(&node2); ans != &node1 {
		t.Fatalf("failed, got %v\n", (*ans).Val)
	}

	if ans := reverseList(&node1); ans != &node2 {
		t.Fatalf("failed, got %v\n", (*ans).Val)
	}

	if ans := reverseList(&node5); ans != &node1 {
		t.Fatalf("failed, got %v\n", (*ans).Val)
	}
}