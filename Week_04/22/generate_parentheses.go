package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	template := make([]byte, 2*n)
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if right == n {
			res = append(res, string(template))
			return
		}
		if left < n {
			template[left+right] = '('
			dfs(left+1, right)
		}
		if right < left {
			template[left+right] = ')'
			dfs(left, right+1)
		}
	}
	if n > 0 {
		dfs(0, 0)
	}
	return res
}
