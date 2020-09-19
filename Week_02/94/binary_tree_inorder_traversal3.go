package main

func inorderTraversal3(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root
			root = root.Left
			pre.Right.Left = nil
		}
	}
	return res
}
