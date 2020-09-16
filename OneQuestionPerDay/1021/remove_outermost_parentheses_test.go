package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		S      string
		target string
	}{
		{"", ""},
		{"()", ""},
		{"(())", "()"},
		{"()()", ""},
		{"(()())", "()()"},
		{"(()())((()))", "()()(())"},
	}

	for _, tt := range tests {
		if ans := removeOuterParentheses(tt.S); ans != tt.target {
			t.Fatalf("S: %s, target: %s, ans: %s\n", tt.S, tt.target, ans)
		}
	}
}
