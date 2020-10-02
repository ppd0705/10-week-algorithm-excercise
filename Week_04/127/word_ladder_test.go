package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		target    int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
	}

	for _, tt := range tests {
		if ans := ladderLength2(tt.beginWord, tt.endWord, tt.wordList); ans != tt.target {
			t.Fatalf("target: %d, ans: %d\n", tt.target, ans)
		}
	}
}
