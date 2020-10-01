package main

func lemonadeChange(bills []int) bool {
	remain := map[int]int{
		5:  0,
		10: 0,
	}
	for _, b := range bills {
		if b == 5 {
			remain[5]++
		} else if b == 10 {
			if remain[5] > 0 {
				remain[5]--
				remain[10]++
			} else {
				return false
			}
		} else {
			if remain[10] > 0 && remain[5] > 0 {
				remain[10]--
				remain[5]--
			} else if remain[5] >2 {
				remain[5] -= 3
			} else {
				return false
			}
 		}
	}
	return true
}
