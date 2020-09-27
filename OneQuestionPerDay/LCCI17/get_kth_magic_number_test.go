package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		k      int
		target int
	}{
		{1, 1},
		{5, 9},
	}

	for _, tt := range tests {
		if ans := getKthMagicNumber2(tt.k); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
