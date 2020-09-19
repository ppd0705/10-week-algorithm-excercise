package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node != nil {
			res = append(res, node.Val)
			preOrder(node.Left)
			preOrder(node.Right)
		}
	}
	preOrder(root)
	return res
}
