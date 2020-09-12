package main

type ListNode struct {
     Val int
     Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	slow, fast := head, head
	for {
		cnt := 1
		for fast != nil && cnt < k {
			fast = fast.Next
			cnt++
		}
		if fast == nil {
			break
		}
		prev.Next = fast
		prev = slow
		next := slow.Next
		prev.Next = fast.Next

		for next != prev.Next {
			next, next.Next, slow = next.Next, slow, next
		}
		slow = prev.Next
		fast = prev.Next
	}
	return dummy.Next
}