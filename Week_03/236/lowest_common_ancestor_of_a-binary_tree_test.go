package main

import "testing"

func Test(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node5.Right = node6
	if ans := lowestCommonAncestor2(node1, node3, node6); ans != node3 {
		t.Fatalf("ans: %v, target: %v\n", ans, node3)
	}
	if ans := lowestCommonAncestor2(node1, node1, node6); ans != node1 {
		t.Fatalf("ans: %v, target: %v\n", ans, node1)
	}

	if ans := lowestCommonAncestor2(node1, node2, node6); ans != node1 {
		t.Fatalf("ans: %v, target: %v\n", ans, node1)
	}

}
