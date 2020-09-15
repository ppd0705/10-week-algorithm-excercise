package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		target bool
	}{
		{"", "", true},
		{"a", "", false},
		{"", "a", false},
		{"abc", "aca", false},
	}

	for _, tt := range tests {
		if ans := isAnagram(tt.s, tt.t); ans != tt.target {
			t.Fatalf("failed. %s %s\n", tt.s, tt.t)
		}
	}
}
