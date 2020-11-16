package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		s      string
		p      string
		target []int
	}{
		{"abab", "a", []int{0, 2}},
		{"abab", "ab", []int{0, 1, 2}},
		{"ababcdba", "ab", []int{0, 1, 2, 6}},
	}

	for _, tt := range tests {
		if ans := findAnagrams2(tt.s, tt.p); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
