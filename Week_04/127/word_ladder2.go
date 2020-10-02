package main

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))

	wordMap := make(map[string]int, 0)
	for i, w := range wordList {
		wordMap[w] = i
	}
	if i, ok := wordMap[endWord]; !ok {
		return 0
	} else {
		endUsed[i] = true
	}
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	count := 1
	n := len(beginWord)
	for len(startQueue) > 0 {
		l := len(startQueue)
		count++
		for i := 0; i < l; i++ {
			chars := []byte(startQueue[i])
			for j := 0; j < n; j++ {
				o := chars[j]
				for c := 'a'; c <= 'z'; c++ {
					chars[j] = byte(c)
					idx, ok := wordMap[string(chars)]
					if !ok || startUsed[idx] {
						continue
					}
					if endUsed[idx] {
						return count
					}
					startUsed[idx] = true
					startQueue = append(startQueue, wordList[idx])
				}
				chars[j] = o
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endUsed) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}
