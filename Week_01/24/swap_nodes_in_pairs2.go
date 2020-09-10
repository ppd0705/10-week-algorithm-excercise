package main

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	for head != nil && head.Next != nil {
		prev.Next = head.Next
		prev = head
		head.Next, head.Next.Next = head.Next.Next, head
		head = head.Next
	}
	return dummy.Next
}
