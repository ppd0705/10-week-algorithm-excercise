package main

func isPowerOfTwo2(n int) bool {
	return n > 0 && n&-n == n
}
