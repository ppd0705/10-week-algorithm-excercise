package main

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
