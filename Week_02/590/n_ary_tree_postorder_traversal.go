package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := make([]int, 0)
	var postOrder func(node *Node)
	postOrder = func(node *Node) {
		if node != nil {
			for i := 0; i < len(node.Children); i++ {
				postOrder(node.Children[i])
			}
			res = append(res, node.Val)
		}
	}
	postOrder(root)
	return res
}
