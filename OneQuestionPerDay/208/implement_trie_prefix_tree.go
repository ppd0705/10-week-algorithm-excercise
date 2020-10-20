package main

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{children: map[byte]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	for i := 0; i < len(word); i++ {
		if _, ok := t.children[word[i]]; !ok {
			t.children[word[i]] = &Trie{children: map[byte]*Trie{}}
		}
		t = t.children[word[i]]
	}
	t.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for i := 0; i < len(word); i++ {
		if _, ok := t.children[word[i]]; !ok {
			return false
		}
		t = t.children[word[i]]
	}
	return t.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		if _, ok := t.children[prefix[i]]; !ok {
			return false
		}
		t = t.children[prefix[i]]
	}
	return true
}


