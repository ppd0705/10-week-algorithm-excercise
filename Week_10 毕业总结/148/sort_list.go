package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tmp.Next = head1
			head1 = head1.Next
		} else {
			tmp.Next = head2
			head2 = head2.Next
		}
		tmp = tmp.Next
	}
	if head1 != nil {
		tmp.Next = head1
	}
	if head2 != nil {
		tmp.Next = head2
	}
	return dummy.Next
}
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(tmp))
}
