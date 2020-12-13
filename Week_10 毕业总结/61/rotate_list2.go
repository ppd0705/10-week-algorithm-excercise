package main

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	node := head
	i := 1
	for ; node.Next != nil; i++ {
		node = node.Next
	}
	k = k % i
	node.Next = head
	for j := 0; j < i-k; j++ {
		node = node.Next
	}
	ret := node.Next
	node.Next = nil
	return ret
}
