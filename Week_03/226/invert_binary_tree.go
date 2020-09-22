package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return root
}
