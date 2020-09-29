package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if len(res) == level {
				res = append(res, []int{queue[i].Val})
			} else {
				res[level] = append(res[level], queue[i].Val)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		level++
	}
	return res
}
