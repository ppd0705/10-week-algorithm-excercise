package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	wordPatternMap := make(map[string]byte, 26)
	wordArr := make([]string, 26)
	for i, word := range words {
		if p, ok := wordPatternMap[words[i]]; !ok {
			if wordArr[pattern[i]-'a'] != "" {
				return false
			}
			wordPatternMap[word] = pattern[i]
			wordArr[pattern[i]-'a'] = word
		} else if pattern[i] != p{
			return false
		}
	}
	return true
}
