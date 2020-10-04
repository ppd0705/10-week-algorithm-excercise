package main

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)
	if idxOf(endWord, wordList) == -1 {
		return ans
	}
	// regex map like d*g: [dog, dwg]
	buckets := map[string][]string{}

	wordList = append(wordList, beginWord)
	for _, w := range wordList {
		letters := []byte(w)
		for i := 0; i < len(letters); i++ {
			o := letters[i]
			letters[i] = byte('*')
			str := string(letters)
			if _, ok := buckets[str]; !ok {
				buckets[str] = []string{w}
			} else {
				buckets[str] = append(buckets[str], w)
			}
			letters[i] = o
		}
	}

	// record step count
	levelMap := map[string]int{beginWord: 1}

	// record pre nodes
	preNodes := map[string][]string{}

	queue := []string{beginWord}
	for len(queue) > 0 {
		curr := queue[0]
		level := levelMap[curr]
		//fmt.Println("queue ", queue)
		//fmt.Println("level ", levelMap)
		//fmt.Println("prev", preNodes)

		// all end word path all ready recorded
		if endLevel, ok := levelMap[endWord]; ok && level+1 > endLevel {
			break
		}

		letters := []byte(curr)
		for i := 0; i < len(letters); i++ {
			o := letters[i]
			letters[i] = byte('*')
			str := string(letters)
			bucket := buckets[str]
			for _, b := range bucket {
				l := levelMap[b]
				if l == 0 {
					// new node
					levelMap[b] = level + 1
					preNodes[b] = []string{curr}
					queue = append(queue, b)
				} else if l == level+1 {
					// new pre node
					preNodes[b] = append(preNodes[b], curr)
				}
			}
			letters[i] = o
		}
		queue = queue[1:]
	}

	var dfs func(word string, level int, solution []string)
	dfs = func(word string, level int, solution []string) {
		solution[level-1] = word
		if level == 1 {
			ans = append(ans, append([]string{}, solution...))
			return
		}
		for _, w := range preNodes[word] {
			dfs(w, level-1, solution)
		}
	}
	if l, ok := levelMap[endWord]; ok {
		solution := make([]string, l)
		dfs(endWord, l, solution)
	}
	return ans
}
