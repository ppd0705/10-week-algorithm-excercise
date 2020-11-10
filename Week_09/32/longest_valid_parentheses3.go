package main

func longestValidParentheses3(s string) int {
	left, right, maxLen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLen = max(maxLen, 2*left)
		} else if left < right {
			left = 0
			right = 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLen = max(maxLen, 2*left)
		} else if right < left {
			left = 0
			right = 0
		}
	}
	return maxLen
}
