package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		target string
	}{
		{"ebebabc", "bab"},
	}
	for _, tt := range tests {
		if ans := longestPalindrome2(tt.s); ans != tt.target {
			t.Fatalf("target: %s, ans: %s\n", tt.target, ans)
		}
	}
}
