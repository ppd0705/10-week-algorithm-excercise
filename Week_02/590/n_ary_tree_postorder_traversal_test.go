package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n1.Children = []*Node{n2, n3, n4}
	n3.Children = []*Node{n5}

	if ans := postorder2(n1); !reflect.DeepEqual(ans, []int{2, 5, 3, 4, 1}) {
		t.Fatalf("ans: %v\n", ans)
	}
}
