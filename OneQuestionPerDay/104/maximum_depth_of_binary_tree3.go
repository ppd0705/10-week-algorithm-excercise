package main

func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	level := 0

	for len(stack) > 0 {
		num := len(stack)
		for i := 0; i < num; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[num:]
		level++
	}
	return level
}
