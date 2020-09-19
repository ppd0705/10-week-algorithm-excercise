package main

func addDigits2(num int) int {
	if num == 0 {
		return 0
	} else {
		return  (num - 1) % 9 + 1
	}
}
