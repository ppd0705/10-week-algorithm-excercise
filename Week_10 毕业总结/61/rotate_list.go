package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	nodeIdx := make(map[int]*ListNode)
	for i := 0; head != nil; i++ {
		nodeIdx[i] = head
		head = head.Next
	}
	k = k % len(nodeIdx)
	if k == 0 {
		return nodeIdx[0]
	}
	nodeIdx[len(nodeIdx)-k-1].Next = nil
	nodeIdx[len(nodeIdx)-1].Next = nodeIdx[0]
	return nodeIdx[len(nodeIdx)-k]
}
