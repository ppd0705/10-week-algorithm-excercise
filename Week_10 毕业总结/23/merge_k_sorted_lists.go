package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	cur := new(ListNode)
	idx := -1
	for {
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if idx == -1 {
					cur = lists[i]
					idx = i
				} else if lists[i].Val < cur.Val {
					cur = lists[i]
					idx = i
				}
			}
		}
		if idx == -1 {
			break
		}
		pre.Next = cur
		lists[idx] = cur.Next
		pre = cur
		idx = -1
	}
	return dummy.Next
}
