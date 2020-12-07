package main

import "testing"

func Test(t  *testing.T) {
	tests := []struct{
		s string
		target int
	}{
		{"abcb", 3},
		{"aabcb", 3},
		{"aaaa", 1},
		{"aaaaert", 4},
	}
	for _, tt  := range tests {
		if ans := lengthOfLongestSubstring(tt.s); ans != tt.target {
			t.Fatalf("s: %s, ans: %d\n", tt.s, ans)
		}
	}
}