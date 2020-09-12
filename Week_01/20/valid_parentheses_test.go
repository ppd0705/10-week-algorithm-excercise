package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s      string
		target bool
	}{
		{"{", false},
		{"[{", false},
		{"[{}", false},
		{"{}", true},
		{"[{]}", false},
	}

	for _, tt := range tests {
		if ans := isValid(tt.s); ans != tt.target {
			t.Fatalf("failed! s: %s\n", tt.s)
		}
	}
}
