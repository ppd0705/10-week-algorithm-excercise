package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1/x
	}
	sub := myPow(x, n/2)
	if n%2 == 0 {
		return sub * sub
	} else {
		return sub * sub * x
	}
}
