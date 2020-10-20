package main

import "testing"

func Test(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	if ans := trie.Search("apple"); !ans {
		t.Fatalf("apple not found")
	}
	if ans := trie.Search("app"); ans {
		t.Fatalf("app is not exit actually")
	}
	if ans := trie.StartsWith("app"); !ans {
		t.Fatalf("prefix search failed")
	}
	trie.Insert("app")
	if ans := trie.Search("app"); !ans {
		t.Fatalf("app not found")
	}

}
