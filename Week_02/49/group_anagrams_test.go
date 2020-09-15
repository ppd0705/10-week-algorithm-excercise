package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		strs   []string
		target [][]string
	}{
		{[]string{""}, [][]string{{""}}},
		{[]string{"a", "ab", "ba"}, [][]string{{"a"}, {"ab", "ba"}}},
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		}}

	for _, tt := range tests {
		if ans := groupAnagrams(tt.strs); !reflect.DeepEqual(ans, tt.target) {
			t.Fatalf("target: %v, ans: %v\n", tt.target, ans)
		}
	}
}
