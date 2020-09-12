package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle0(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	for head != nil {
		if seen[head] {
			return true
		} else {
			seen[head] = true
		}
		head = head.Next
	}
	return false
}
