package main

import (
	"testing"
)

func Test(t *testing.T) {
	stack := Constructor()
	stack.Push(3)
	if stack.GetMin() != 3 {
		t.Fatalf("epected GetMin() 3, got %d\n", stack.GetMin())
	}
	stack.Push(1)
	stack.Push(-1)
	stack.Push(-1)
	stack.Push(2)
	if stack.GetMin() != -1 {
		t.Fatalf("epected GetMin() -1, got %d\n", stack.GetMin())
	}
	if stack.Top() != 2 {
		t.Fatalf("epected Top 2, got %d\n", stack.Top())
	}
	stack.Pop()
	stack.Pop()
	if stack.GetMin() != -1 {
		t.Fatalf("epected GetMin() -1, got %d\n", stack.GetMin())
	}
	stack.Pop()
	if stack.GetMin() != 1 {
		t.Fatalf("epected GetMin() 1, got %d\n", stack.GetMin())
	}
	stack.Pop()
	if stack.GetMin() != 3 {
		t.Fatalf("epected GetMin() 3, got %d\n", stack.GetMin())
	}
}