package main

func countSubstrings2(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < 2*n-1; i++ {
		l := i /2
		r := l + i % 2
		for l >=0 && r <n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
