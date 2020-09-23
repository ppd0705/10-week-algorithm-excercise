package main

import "math"

func isValidBST2(root *TreeNode) bool {
	pre := math.MinInt64

	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		root = root.Right
	}
	return true
}
