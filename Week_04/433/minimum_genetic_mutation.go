package main


func minMutation(start string, end string, bank []string) int {
	visited := make(map[string]bool, len(bank))
	for _, s := range bank {
		visited[s] = false
	}
	if _, ok := visited[end]; !ok {
		return -1
	}
	mutationMap := map[byte]string{
		'A': "CGT",
		'C': "AGT",
		'G': "CAT",
		'T': "CGA",
	}
	queue := []string{start}
	num := 0
	for len(queue) > 0 {
		n := len(queue)
		num++
		for i:=0;i <n;i++ {
			s := queue[i]
			for j := 0; j < len(s);j++ {
				for _, b := range mutationMap[s[j]] {
					ss := s[:j] + string(b) + s[j+1:]
					if ss == end {
						return num
					}
					if used, ok := visited[ss]; ok && !used {
						visited[ss] = true
						queue = append(queue, ss)
					}
				}
			}
		}
		queue = queue[n:]
	}
	return -1
}
