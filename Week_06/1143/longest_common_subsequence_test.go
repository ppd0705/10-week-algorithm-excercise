package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		text1  string
		text2  string
		target int
	}{
		{"abcba", "abcbcba", 5},
		{"ace", "abcde", 3},
	}

	for _, tt := range tests {
		if ans := longestCommonSubsequence3(tt.text1, tt.text2); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
