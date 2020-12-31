package main

import (
	"container/heap"
	"fmt"
)

type ListNodes []*ListNode

func (l ListNodes) Len() int {
	return len(l)
}

func (l ListNodes) Less(i, j int) bool {
	return l[i].Val <= l[j].Val
}

func (l ListNodes) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListNodes) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *ListNodes) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

var _ heap.Interface = (*ListNodes)(nil)

func mergeKLists2(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	hp := &ListNodes{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(hp, lists[i])
		}
	}
	for len(*hp) > 0 {
		top := heap.Pop(hp).(*ListNode)
		pre.Next = top
		pre = pre.Next
		if top.Next != nil {
			heap.Push(hp, top.Next)
		}
	}
	return dummy.Next
}

func printHP(hp ListNodes) {
	fmt.Printf("\nhp: len: %d \n", len(hp))
	for i := 0; i < len(hp); i++ {
		fmt.Printf("node %d: %v", i, (hp)[i])
	}
}

func printNode(node *ListNode) {
	fmt.Printf("\nnodes: ")
	for node != nil {
		fmt.Printf(" %v ->", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}
