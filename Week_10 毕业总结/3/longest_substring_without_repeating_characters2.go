package main

func lengthOfLongestSubstring2(s string) int {
	left, res := -1, 0
	position := make(map[byte]int)
	for right := 0; right < len(s);right++ {
		if p, ok := position[s[right]]; ok && p > left {
			left = p
		} else {
			res = max(res, right-left)
		}
		position[s[right]] = right
	}
	return res
}