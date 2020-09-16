package main

import "strings"

func removeOuterParentheses(S string) string {
	var strs []string
	start := 0
	counter := 0
	for i, c := range S {
		if c == '(' {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			strs = append(strs, S[start+1:i])
			start = i + 1
		}
	}
	return strings.Join(strs, "")
}
