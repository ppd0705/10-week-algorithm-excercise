package main

func largestRectangleArea1(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	n := len(heights)
	leftIdx := make([]int, n)
	rightIdx := make([]int, n)
	leftIdx[0] = -1
	rightIdx[n-1] = n

	for i := 1; i < n; i++ {
		l := i - 1
		for l >= 0 && heights[l] >= heights[i] {
			l = leftIdx[l]
		}
		leftIdx[i] = l
	}
	for i := n - 2; i >= 0; i-- {
		r := i + 1
		for r < n && heights[r] >= heights[i] {
			r = rightIdx[r]
		}
		rightIdx[i] = r
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, heights[i]*(rightIdx[i]-leftIdx[i]-1))
	}
	return res
}
