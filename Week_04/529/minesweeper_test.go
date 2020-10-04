package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		board [][]byte
		click []int
	}{
		{[][]byte{{'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'M', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}},
			[]int{3, 0}},
		{[][]byte{{'B', '1', 'E', '1', 'B'}, {'B', '1', 'M', '1', 'B'}, {'B', '1', '1', '1', 'B'}, {'B', 'B', 'B', 'B', 'B'}}, []int{0, 2}},
	}

	for _, tt := range tests {
		for i := 0; i < len(tt.board); i++ {
			fmt.Printf("%c\n", tt.board[i])
		}
		println()
		updateBoard2(tt.board, tt.click)
		for i := 0; i < len(tt.board); i++ {
			fmt.Printf("%c\n", tt.board[i])
		}
	}

}
