package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)

	var levelOrder func(node *Node, level int)
	levelOrder = func(node *Node, level int) {
		if node != nil {
			if len(res) == level {
				res = append(res, []int{node.Val})
			} else {
				res[level] = append(res[level], node.Val)
			}
			for i := 0; i < len(node.Children); i++ {
				levelOrder(node.Children[i], level+1)
			}
		}
	}
	levelOrder(root, 0)
	return res
}
