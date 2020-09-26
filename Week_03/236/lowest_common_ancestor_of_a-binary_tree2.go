package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode, 0)
	visited := make(map[*TreeNode]bool, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[p] = true
		p = parent[p]
	}
	for q != nil {
		if visited[q] {
			return q
		}
		q = parent[q]
	}
	return nil
}
