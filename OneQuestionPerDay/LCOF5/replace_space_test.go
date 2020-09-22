package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		s      string
		target string
	}{
		{"abc", "abc"},
		{"a bc", "a%20bc"},
		{" a bc", "%20a%20bc"},
		{" a bc ", "%20a%20bc%20"},
	}

	for _, tt := range tests {
		if ans := replaceSpace(tt.s); ans != tt.target {
			t.Fatalf("target: %s, ans: %s\n", tt.target, ans)
		}
	}
}
