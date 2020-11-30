package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := make([]string, 0)
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
 		} else if  cur != ']' {
 			stack = append(stack, string(cur))
 			ptr++
		} else {
			ptr++
			sub := make([]string, 0)

			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// reverse
			for i := 0; i < len(sub)/2;i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stack = stack[:len(stack)-1]
			repeatTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack[len(stack)-1] = strings.Repeat(getString(sub), repeatTime)
		}
	}
	return getString(stack)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string{
	return strings.Join(v, "")
}