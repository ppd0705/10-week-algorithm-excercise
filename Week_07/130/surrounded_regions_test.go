package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		board  [][]byte
		target [][]byte
	}{
		{[][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}, [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		},
	}
	for _, tt := range tests {
		solve2(tt.board)
		if !reflect.DeepEqual(tt.board, tt.target) {
			t.Fatalf("board:\n %v\n, target:\n %v\n", tt.board, tt.target)
		}
	}
}
