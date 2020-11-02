package main

func findLength2(A []int, B []int) int {
	m, n := len(A), len(B)
	if m == 0 || n == 0 {
		return 0
	}
	res := 0

	for i := 0; i < m; i++ {
		l := min(m-i, n)
		res = max(res, maxLen(A, i, B, 0, l))
	}
	for i := 0; i < n; i++ {
		l := min(n-i, m)
		res = max(res, maxLen(A, 0, B, i, l))
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func maxLen(a []int, i int, b []int, j int, m int) int {
	res := 0
	k := 0
	for l := 0; l < m; l++ {
		if a[i+l] == b[j+l] {
			k++
			res = max(res, k)
		} else {
			k = 0
		}
	}
	return res
}
