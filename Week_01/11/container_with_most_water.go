package main

func maxArea(height []int) int {
	i, j, res, v := 0, len(height)-1, 0, 0
	for i < j {
		if height[i] <= height[j] {
			v = height[i] * (j - i)
			i++
		} else {
			v = height[j] * (j - i)
			j--
		}
		if v > res {
			res = v
		}
	}
	return res
}
