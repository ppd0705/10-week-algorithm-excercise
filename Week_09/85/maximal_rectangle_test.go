package main

import "testing"

func Test(t *testing.T) {
	matrix := [][]byte{
		//{'1', '0', '1', '0', '0'},
		//{'1', '0', '1', '1', '1'},
		//{'1', '1', '1', '1', '1'},
		//{'1', '0', '0', '1', '0'},
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
	}
	target := 6
	if ans := maximalRectangle(matrix); ans != target {
		t.Fatalf("target: %d, ans: %d\n", target, ans)
	}

}
