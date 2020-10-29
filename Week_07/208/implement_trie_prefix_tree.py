class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.nodes = {}
        self.end_mark = "#"

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        nodes = self.nodes

        for w in word:
            nodes = nodes.setdefault(w, {})
        nodes[self.end_mark] = {}

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        nodes = self.nodes
        for w in word:
            if w not in nodes:
                return False
            nodes = nodes[w]
        return self.end_mark in nodes

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        nodes = self.nodes
        for w in prefix:
            if w not in nodes:
                return False
            nodes = nodes[w]
        return True
