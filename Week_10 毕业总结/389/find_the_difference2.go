package main

func findTheDifference2(s string, t string) byte {
	var ret byte
	for i := 0; i < len(s); i++ {
		ret ^= s[i]
	}
	for i := 0; i < len(t);i++ {
		ret ^= t[i]
	}
	return ret
}