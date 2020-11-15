package main

func reverseWords(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(bytes) && bytes[j] != ' ' {
			j++
		}
		reverseHelper(bytes, i, j-1)
		i = j + 1
	}
	return string(bytes)
}

func reverseHelper(bytes []byte, start, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}
