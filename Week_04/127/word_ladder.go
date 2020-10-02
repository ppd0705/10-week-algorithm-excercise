package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	isUsed := make(map[string]bool, 0)
	for _, w := range wordList {
		isUsed[w] = false
	}
	if _, ok := isUsed[endWord]; !ok {
		return 0
	}
	count := 1
	n := len(beginWord)
	queue := []string{beginWord}
	for len(queue) > 0 {
		length := len(queue)
		count++
		for k := 0; k < length; k++ {
			node := queue[k]
			for i := 0; i < n; i++ {
				for j := uint8(0); j < 26; j++ {
					if node[i] == 'a'+j {
						continue
					}
					s := node[:i] + string('a'+j) + node[i+1:]
					if s == endWord {
						return count
					}
					if used, ok := isUsed[s]; ok && !used {
						isUsed[s] = true
						queue = append(queue, s)
					}
				}
			}
		}
		queue = queue[length:]
	}
	return 0
}
