package main

func preorder2(root *Node) []int {
	res := make([]int, 0)
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return res
}
