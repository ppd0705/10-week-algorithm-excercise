package main

import (
	"log"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		n      int
		target []string
	}{
		{1, []string{"1"}},
		{3, []string{"1", "2", "Fizz"}},
		{15, []string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		}},
	}

	for _, tt := range tests {
		if ans := fizzBuzz(tt.n); !reflect.DeepEqual(ans, tt.target) {
			log.Fatalf("n: %d, ans: %v\n", tt.n, ans)
		}
	}
}
