package main

func getLength(node *ListNode) int {
	n := 0
	for node != nil {
		node = node.Next
		n++
	}
	return n
}

func split(head *ListNode, step int) *ListNode {
	if head == nil {
		return nil
	}
	for i := 1; i < step && head.Next != nil; i++ {
		head = head.Next
	}
	right := head.Next
	head.Next = nil
	return right
}

func sortList2(head *ListNode) *ListNode {
	length := getLength(head)
	dummy := &ListNode{Next: head}
	for step := 1; step < length; step = step*2 {
		prev, cur := dummy, dummy.Next
		for cur != nil {
			h1 := cur
			h2 := split(h1, step)
			cur = split(h2, step)
			prev.Next = merge(h1, h2)
			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}
	return dummy.Next
}
