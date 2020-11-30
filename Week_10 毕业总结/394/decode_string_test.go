package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		target string
	}{
		{"abc", "abc"},
		{"2[ab]c", "ababc"},
		{"2[a2[ab]]c", "aababaababc"},
	}

	for _, tt := range tests {
		if ans := decodeString(tt.s); ans != tt.target {
			t.Fatalf("ans: %s, target: %s\n", ans, tt.target)
		}
	}
}
