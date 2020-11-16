package main

func isLetter(c byte) bool {
	if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

func reverseOnlyLetters(S string) string {
	bytes := []byte(S)
	left, right := 0, len(bytes)-1
	for left < right {
		if !isLetter(bytes[left]) {
			left++
		} else if !isLetter(bytes[right]) {
			right--
		} else {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}
