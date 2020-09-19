package main

func addDigits(num int) int {
	for num > 9 {
		sum := 0
		for num > 9 {
			sum += num / 10
			num %= 10
		}
		num += sum
	}
	return num
}
