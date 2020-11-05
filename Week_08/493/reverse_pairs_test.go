package main

import "testing"

func Test(t *testing.T) {
	if ans := reversePairs([]int{2, 4, 3, 5, 1}); ans != 3 {
		t.Fatalf("failed, ans: %d\n", ans)
	}
}
