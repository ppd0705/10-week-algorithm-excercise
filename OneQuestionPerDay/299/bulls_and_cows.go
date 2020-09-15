package main

import "strconv"

func getHint(secret string, guess string) string {
	secretCount := make([]int, 10)
	guessCount := make([]int, 10)
	bulls := 0
	cows := 0

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretCount[secret[i]-'0']++
			guessCount[guess[i]-'0']++
		}
	}
	for i, v := range secretCount {
			cows += min(v, guessCount[i])
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
