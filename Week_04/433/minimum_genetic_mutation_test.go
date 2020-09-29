package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		start  string
		end    string
		bank   []string
		target int
	}{
		{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
		{"GAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, -1},
	}

	for _, tt := range tests {
		if ans := minMutation(tt.start, tt.end, tt.bank); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
