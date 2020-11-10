package main

func longestValidParentheses2(s string) int {
	maxLen := 0
	stack := []int{-1}
	for i := 0; i < len(s);i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}
