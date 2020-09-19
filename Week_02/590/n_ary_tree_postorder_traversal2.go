package main

func postorder2(root *Node) []int {
	res := make([]int, 0)
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			for i := 0; i < len(node.Children); i++ {
				stack = append(stack, node.Children[i])
			}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
