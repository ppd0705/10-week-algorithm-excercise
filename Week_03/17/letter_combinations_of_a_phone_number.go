package main

func letterCombinations(digits string) []string {
	numberLetterMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	base := make([]byte, len(digits))
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(digits) {
			res = append(res, string(base))
			return
		}
		for _, b := range numberLetterMap[digits[idx]] {
			base[idx] = b
			dfs(idx + 1)
		}
	}
	dfs(0)
	return res
}
