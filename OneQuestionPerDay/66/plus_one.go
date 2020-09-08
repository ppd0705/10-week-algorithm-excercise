package main

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			carry = false
			break
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
