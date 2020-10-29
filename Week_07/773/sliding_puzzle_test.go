package main

import "testing"

func Test(t *testing.T){
	tests := []struct{
		board [][]int
		target int
	}{
		{[][]int{{1,2,3}, {5,4,0}}, -1},
	}
	for _, tt := range tests {
		if ans := slidingPuzzle(tt.board); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}

}
