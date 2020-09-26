package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	vals := make([]string, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			vals = append(vals, "#")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	return strings.Join(vals, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	if vals[0] == "#" {
		return nil
	}
	idx := 0
	val, _ := strconv.Atoi(vals[idx])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	var node, left, right *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			left = &TreeNode{Val: val}
			node.Left = left
			queue = append(queue, left)
		}
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			right = &TreeNode{Val: val}
			node.Right = right
			queue = append(queue, right)
		}

	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
