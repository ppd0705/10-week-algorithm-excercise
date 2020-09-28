package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(arr []int, idx int)
	dfs = func(arr []int, idx int) {
		res = append(res, append([]int{}, arr...))
		for i := idx; i < len(nums); i++ {
			dfs(append(arr, nums[i]), i+1)
		}
	}
	dfs([]int{}, 0)
	return res
}
