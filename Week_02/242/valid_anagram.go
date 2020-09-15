package main

func isAnagram(s string, t string) bool {
	counter := make([]int, 26)

	for _, c := range s {
		counter[c-'a']++
	}
	for _, c := range t {
		counter[c-'a']--
	}

	for _, i := range counter {
		if i != 0 {
			return false
		}
	}
	return true
}