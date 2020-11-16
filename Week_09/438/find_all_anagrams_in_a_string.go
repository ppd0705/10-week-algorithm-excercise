package main

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	base := countLetters(p)
	for i := 0; i <= len(s)-len(p); i++ {
		if isAnagram(base, countLetters(s[i:i+len(p)])) {
			res = append(res, i)
		}
	}
	return res
}

func countLetters(s string) []int {
	res := make([]int, 26)
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	return res
}

func isAnagram(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
