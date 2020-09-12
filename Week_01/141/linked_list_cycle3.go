package main


func hasCycle3(head *ListNode) bool {
	if head ==  nil {
		return false
	}
	prev, cur := head, head.Next
	for cur != nil {
		if cur == head {
			return true
		}
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return false
}
