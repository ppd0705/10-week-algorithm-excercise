package main

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			leftFlag, rightFlag := true, true
			for i, j := left+1, right; i < j; i, j= i+1, j-1 {
				if s[i] != s[j] {
					leftFlag = false
					break
				}
			}
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					rightFlag = false
					break
				}
			}
			return leftFlag || rightFlag
		}
	}
	return true
}
