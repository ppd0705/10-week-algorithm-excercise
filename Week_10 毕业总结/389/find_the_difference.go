package main

func findTheDifference(s string, t string) byte {
	counter  := [26]int{}
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']--
		counter[t[i]-'a']++
	}
	counter[t[len(t)-1]-'a']++
	for i := 0; i < 26;i++ {
		if counter[i] == 1 {
			return byte('a'+i)
		}
	}
	return 0
}