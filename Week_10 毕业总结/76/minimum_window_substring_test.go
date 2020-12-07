package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		target string
	}{
		{"abc", "e", ""},
		{"abc", "a", "a"},
		{"abdeecaedb", "abc", "a"},
	}

	for _, tt := range tests {
		if ans := minWindow(tt.s, tt.t); ans != tt.target {
			t.Fatalf("ans: %s, target: %s\n", ans, tt.target)
		}
	}
}
