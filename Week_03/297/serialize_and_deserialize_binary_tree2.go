package main

import (
	"strconv"
	"strings"
)

type Codec2 struct {
}

func Constructor2() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec2) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec2) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}

func dfs(valsPtr *[]string) *TreeNode{
	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]
	if val == "#" {
		return nil
	}
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{Val: valInt}
	node.Left = dfs(valsPtr)
	node.Right = dfs(valsPtr)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
