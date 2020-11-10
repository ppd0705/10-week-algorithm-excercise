package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	elements := map[int]bool{head.Val: true}
	prev, cur := head, head.Next
	for cur != nil {
		if elements[cur.Val] {
			prev.Next, cur = cur.Next, cur.Next
		} else {
			elements[cur.Val] = true
			prev, cur = cur, cur.Next
		}
	}
	return head
}
