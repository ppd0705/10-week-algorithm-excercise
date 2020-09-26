package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil{
		return 1
	} else if node.Left != nil && node.Right != nil {
		return 1 + min(minDepth(node.Left), minDepth(node.Right))
	} else if node.Left != nil {
		return 1 + minDepth(node.Left)
	} else {
		return 1 + minDepth(node.Right)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}