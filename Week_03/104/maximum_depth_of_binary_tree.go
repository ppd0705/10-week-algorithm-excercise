package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(dfs(node.Left), dfs(node.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return dfs(root)
}
