package main

func reverseList0(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
	headNew := head
	for head.Next != nil {
		next := head.Next
		head.Next, next.Next = next.Next, headNew
		headNew = next
	}
	return headNew
}
