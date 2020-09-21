package main

func generateParenthesis(n int) []string {
	var res []string
	solution := make([]byte, n*2)
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if right == n {
			res = append(res, string(solution))
			return
		}
		if left < n {
			solution[left+right] = '('
			dfs(left+1, right)
		}
		if right < left {
			solution[left+right] = ')'
			dfs(left, right+1)
		}
	}
	dfs(0, 0)
	return res
}
