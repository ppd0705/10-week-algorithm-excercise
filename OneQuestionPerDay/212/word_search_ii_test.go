package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		board  [][]byte
		words  []string
		target []string
	}{
		{[][]byte{{'a'}, {'a'}}, []string{"aa"}, []string{"aa"}},
	}

	for _, tt := range tests {
		if ans := findWords(tt.board, tt.words); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
