package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		matrix [][]int
		n      int
		target bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			}, 3, true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			}, 2, false,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			}, 0, false,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			}, 60, false,
		},
	}

	for _, tt := range tests {
		if ans := searchMatrix(tt.matrix, tt.n); ans != tt.target {
			t.Fatalf("ans: %T\n", ans)
		}
	}
}
