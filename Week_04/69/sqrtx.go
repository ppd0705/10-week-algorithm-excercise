package main

func mySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == x {
			return mid
		} else if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
