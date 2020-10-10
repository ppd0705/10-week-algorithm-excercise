package main

type point struct {
	x, y int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp func(i, j int) int
	cache := map[point]int{}
	dp = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		p := point{i, j}
		if v, ok := cache[p]; ok {
			return v
		}
		if text1[i] == text2[j] {
			cache[p] = dp(i-1, j-1) + 1
		} else {
			cache[p] = max(dp(i-1, j), dp(i, j-1))
		}
		return cache[p]
	}
	return dp(len(text1)-1, len(text2)-1)
}
