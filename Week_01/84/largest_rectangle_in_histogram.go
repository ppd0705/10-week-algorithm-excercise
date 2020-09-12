package main

func largestRectangleArea0(heights []int) int {
	res := 0

	for i := 0; i < len(heights); i++ {
		l, r := i, i
		for l > 0 && heights[l-1] >= heights[i] {
			l--
		}
		for r < len(heights)-1 && heights[r+1] >= heights[i] {
			r++
		}
		res = max(res, heights[i]*(r-l+1))
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
