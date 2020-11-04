package main

func isPowerOfThree2(n int) bool {
	i := 1
	for i < n {
		i  *= 3
	}
	return n == i
}