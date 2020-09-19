package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		num    int
		target int
	}{
		{3, 3},
		{38, 2},
		{19, 1},
		{138, 3},
		{888, 6},
	}

	for _, tt := range tests {
		if ans := addDigits(tt.num); ans != tt.target {
			t.Fatalf("ans: %d, target: %d\n", tt.num, tt.target)
		}
	}
}
