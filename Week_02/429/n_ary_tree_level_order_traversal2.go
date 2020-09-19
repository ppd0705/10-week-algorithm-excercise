package main

func levelOrder2(root *Node) [][]int {
	res := make([][]int, 0)
	stack := []*Node{root}
	level := 0
	for len(stack) > 0 {
		num := len(stack)
		for i := 0; i < num; i++ {
			if stack[i] != nil {
				if len(res) == level {
					res= append(res, []int{stack[i].Val})
				} else {
					res[level] = append(res[level], stack[i].Val)
				}
				for j := 0; j < len(stack[i].Children); j++ {
					stack = append(stack, stack[i].Children[j])
				}
			}
		}
		stack = stack[num:]
		level++
	}
	return res
}
