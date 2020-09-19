package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(node *TreeNode, res []int) []int {
	if node != nil {
		res = inorder(node.Left, res)
		res = append(res, node.Val)
		res = inorder(node.Right, res)
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	return inorder(root, []int{})
}
