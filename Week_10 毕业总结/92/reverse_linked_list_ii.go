package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 1; i < n && head != nil && head.Next != nil; i++ {
		if i < m {
			prev = head
			head = head.Next
		} else {
			nxt := head.Next
			head.Next = nxt.Next
			nxt.Next = prev.Next
			prev.Next = nxt
		}
	}
	return dummy.Next
}
