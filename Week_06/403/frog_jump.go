package main

type crossParam struct {
	start int
	jump  int
}

func canCross(stones []int) bool {
	stoneSet := make(map[int]bool, 0)
	for _, s := range stones {
		stoneSet[s] = true
	}
	memo := make(map[crossParam]bool, 0)
	var helper func(i, k int) bool
	helper = func(i, k int) bool {
		if i == stones[len(stones)-1] {
			return true
		}
		if v, ok := memo[crossParam{i, k}]; ok {
			return v
		}
		res := false
		for j := k - 1; j <= k+1; j++ {
			if j > 0 && stoneSet[i+j] && helper(i+j, j) {
				res = true
				break
			}
		}
		memo[crossParam{i, k}] = res
		return res
	}
	return helper(0, 0)
}

