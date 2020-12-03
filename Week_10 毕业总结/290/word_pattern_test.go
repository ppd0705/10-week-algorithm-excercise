package main

import "testing"

func  Test(t *testing.T) {
	tests  := []struct {
		patten string
		words string
		target bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}

	for _, tt := range tests {
		if ans := wordPattern(tt.patten, tt.words); ans != tt.target {
			t.Fatalf("sample: %v\n", tt)
		}
	}
}