package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		secret string
		guess  string
		target string
	}{
		{"", "", "0A0B"},
		{"1807", "7810", "1A3B"},
		{"1123", "0111", "1A1B"},
	}

	for _, tt := range tests {
		if ans := getHint(tt.secret, tt.guess); ans != tt.target {
			t.Fatalf("target: %s, ans: %s\n", tt.target, ans)
		}
	}
}
