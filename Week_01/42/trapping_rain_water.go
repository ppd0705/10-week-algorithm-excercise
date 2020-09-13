package main

func trap(height []int) int {
	stack := make([]int, 0)
	res := 0

	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			res += (i - left - 1) * (min(h, height[left]) - height[top])
		}
		stack = append(stack, i)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
