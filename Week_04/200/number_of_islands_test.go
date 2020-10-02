package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		target int
	}{
		{[][]byte{
			{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'},
		}, 1},
	}

	for _, tt := range tests {
		if ans := numIslands(tt.grid); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
