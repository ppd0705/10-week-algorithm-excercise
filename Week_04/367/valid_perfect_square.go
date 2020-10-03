package main

func isPerfectSquare(num int) bool {
	l, r := 0, num

	for l <= r {
		m := (l + r) / 2
		square := m * m
		if square == num {
			return true
		} else if square < num {
			l = m + 1
		} else {
			r = m - 1
		}

	}
	return false
}
