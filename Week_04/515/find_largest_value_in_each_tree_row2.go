package main

func largestValues2(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, node.Val)
		} else if node.Val > res[level] {
			res[level] = node.Val
		}

		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	return res
}
