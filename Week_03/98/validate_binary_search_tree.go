package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	return node.Val > left && node.Val < right &&
		dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
}
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}
