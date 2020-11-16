package main

func isAlphanumeric(b byte) bool {
	if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
		} else if !isAlphanumeric(s[right]) {
			right--
		} else {
			if s[left] == s[right] || s[left] >= 'a' && s[left] <= 'z' && s[left]-s[right] == 32 || s[right] >= 'a' && s[right] <= 'z' && s[right]-s[left] == 32 {
				left++
				right--
			} else {
				return false
			}
		}
	}
	return true
}
