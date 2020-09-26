package main

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	dep := 0
	for len(queue) > 0 {
		dep++
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return dep
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]

	}
	return -1 // never reached
}

