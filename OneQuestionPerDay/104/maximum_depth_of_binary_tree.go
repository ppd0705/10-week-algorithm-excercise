package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	n := -1
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			if level >= n {
				n = level + 1
			}
			return
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, n)
	return n
}
