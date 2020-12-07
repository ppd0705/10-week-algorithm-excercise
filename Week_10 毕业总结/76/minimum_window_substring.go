package main

import (
	"math"
)

func check(a, b map[byte]int) bool {
	for k, v := range a {
		if b[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	left, right, start, minLen := 0, 0, -1, math.MaxInt32
	target, cur := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}
	for ; right < len(s); right++ {
		if _, ok := target[s[right]]; ok {
			cur[s[right]]++
		}
		for check(target, cur) && left <= right {
			if right-left+1 < minLen {
				start, minLen = left, right-left+1
			}
			if cur[s[left]] > 0 {
				cur[s[left]]--
			}
			left++
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+minLen]
}
