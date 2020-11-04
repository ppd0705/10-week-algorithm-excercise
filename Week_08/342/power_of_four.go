package main

func isPowerOfFour(num int) bool {
	for num > 0 && num % 4 == 0 {
		num /= 4
	}
	return num == 1
}