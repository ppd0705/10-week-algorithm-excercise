package main


func lengthOfLongestSubstring(s string) int {
	right, res := 0, 0
	set := make(map[byte]bool)
	for i := 0; i < len(s);i++ {
		if i != 0 {
			delete(set, s[i-1]) // sliding 1 step
		}
		for right < len(s) && !set[s[right]] {
			set[s[right]] = true
			right++
		}
		res = max(res, right-i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}