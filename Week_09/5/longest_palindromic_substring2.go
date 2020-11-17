package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func centerSpread(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1
}

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		l1 := centerSpread(s, i, i)
		l2 := centerSpread(s, i, i+1)
		l := max(l1, l2)
		if l > end-start+1 {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}
