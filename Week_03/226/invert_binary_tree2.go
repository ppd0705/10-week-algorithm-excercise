package main

func invertTree2(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			queue = append(queue, node.Left, node.Right)
		}
	}
	return root
}
