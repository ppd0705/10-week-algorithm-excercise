package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		target [][]int
	}{
		{4, 2, [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4}}},
		{1, 1, [][]int{{1}}},
	}

	for _, tt := range tests {
		if ans := combine(tt.n, tt.k); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
