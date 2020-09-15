package main

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	indexMap := make(map[string]int, 0)
	for _, s := range strs {
		ss := sortString(s)
		if idx, ok := indexMap[ss]; !ok {
			indexMap[ss] = len(indexMap)
			res = append(res, []string{s})
		} else {
			res[idx] = append(res[idx], s)
		}
	}
	return res
}

func sortString(str string) string {
	sl := make([]rune, 26)

	for _, c := range str {
		sl[c-'a']++
	}
	return string(sl)
}
