package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]
	idx := 0
out:
	for idx < len(ans) {
		for i := 1; i < len(strs); i++ {
			if idx >= len(strs[i]) || strs[i][idx] != ans[idx] {
				break out
			}
		}
		idx++
	}
	return ans[:idx]
}
