package main

func largestRectangleArea(heights []int) int {
	res := 0
	var stack []int
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	for i, h := range heights {
		for len(stack) > 1 && h < heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			res = max(res, heights[idx]*(i-stack[len(stack)-2]-1))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
