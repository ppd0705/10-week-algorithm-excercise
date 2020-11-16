package main

func findAnagrams2(s string, p string) []int {
	res := make([]int, 0)

	counter := make([]int, 26)
	for i := 0; i < len(p); i++ {
		counter[p[i]-'a']++
	}
	tmpCounter := make([]int, 26)
	for left, right := 0, 0; right < len(s); right++ {
		tmpCounter[s[right]-'a']++
		for tmpCounter[s[right]-'a'] > counter[s[right]-'a'] {
			tmpCounter[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) {
			res = append(res, left)
		}
	}
	return res
}
