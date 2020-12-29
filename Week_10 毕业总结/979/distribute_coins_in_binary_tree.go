package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	step := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		step += abs(left) + abs(right)
		return left + right + node.Val - 1
	}
	dfs(root)
	return step
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
