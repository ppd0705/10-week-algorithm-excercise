package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	var preOrder func(node *Node)
	preOrder = func(node *Node) {
		if node != nil {
			res = append(res, node.Val)
			for i := 0; i < len(node.Children); i++ {
				preOrder(node.Children[i])
			}
		}
	}
	preOrder(root)
	return res
}
