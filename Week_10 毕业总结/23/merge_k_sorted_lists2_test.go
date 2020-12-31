package main

import "testing"

func Test(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	node1.Next = node2

	node3.Next = node5

	tests := []struct {
		lists  []*ListNode
		target *ListNode
	}{
		{[]*ListNode{node1, node4, node6, node3}, node1},
	}

	for _, tt := range tests {
		if ans := mergeKLists2(tt.lists); ans != tt.target {
			t.Fatalf("ans:  %v, target: %v", ans, tt.target)
		} else{
			printNode(ans)
		}
	}
}
