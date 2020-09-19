package main

func preorderTraversal3(root *TreeNode) []int {
	res := make([]int, 0)

	for root != nil {
		res = append(res, root.Val)
		if root.Left == nil {
			root = root.Right
		} else {
			pre := root.Left

			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root.Right
			root = root.Left
		}
	}
	return res
}
