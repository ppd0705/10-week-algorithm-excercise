package main

import "testing"

func Test(t *testing.T) {
	q := Constructor(3)
	if !q.IsEmpty() {
		t.Fatalf("failed")
	}
	if q.GetFront() != -1 {
		t.Fatalf("failed")
	}
	q.InsertLast(3)
	if ans := q.GetFront(); ans != 3 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	if ans := q.GetRear(); ans != 3 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	q.InsertLast(8)
	if ans := q.GetFront(); ans != 3 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	if ans := q.GetRear(); ans != 8 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	q.InsertLast(20)
	if !q.IsFull() {
		t.Fatalf("failed")
	}
	if q.InsertLast(30) {
		t.Fatalf("failed\n")
	}
	q.DeleteFront()
	if ans := q.GetFront(); ans != 8 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	if ans := q.GetRear(); ans != 20 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	q.DeleteLast()
	if ans := q.GetFront(); ans != 8 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	if ans := q.GetRear(); ans != 8 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	q.DeleteLast()
	if ans := q.GetFront(); ans != -1 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	if ans := q.GetRear(); ans != -1 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
	if !q.IsEmpty() {
		t.Fatalf("failed")
	}
}
