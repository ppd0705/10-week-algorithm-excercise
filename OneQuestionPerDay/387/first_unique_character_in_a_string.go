package main

func firstUniqChar(s string) int {
	arr := [26]int{}

	for i := 0;i < len(s);i++ {
		arr[s[i]-'a'] = i
	}
	for i := 0 ; i < len(s);i++ {
		if arr[s[i]-'a'] == i {
			return i
		}
		arr[s[i]-'a'] = -1
	}
	return -1
}