package main

import "math"

func myAtoi(str string) int {
	res := 0
	sign := 1
	idx := 0

	for idx < len(str) && str[idx] == ' ' {
		idx++
	}
	if idx == len(str) {
		return 0
	}
	if str[idx] == '-' {
		idx++
		sign = -1
	} else if str[idx] == '+' {
		idx++
	}
	for idx < len(str) && str[idx] >= '0' && str[idx] <= '9' {
		if sign == 1 {
			if res > math.MaxInt32 / 10 || res == math.MaxInt32/10 && str[idx] >= '7'{
				return math.MaxInt32
			}
			res = res*10 + int(str[idx]-'0')
		} else {
			if res < math.MinInt32 / 10 || res == math.MinInt32/10 && str[idx] >= '8'{
				return math.MinInt32
			}
			res = res*10 - int(str[idx]-'0')
		}
		idx++
	}
	return res
}
