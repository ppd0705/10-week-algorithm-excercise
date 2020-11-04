package main

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	n := len(bytes)
	for i:=0; i < n; {
		if n - i < k {
			reverseHelper(bytes, i, n-1)
		} else {
			reverseHelper(bytes, i, i+k-1)
		}
		i  += 2*k
	}
	return string(bytes)
}

func reverseHelper(bytes []byte, start int, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}