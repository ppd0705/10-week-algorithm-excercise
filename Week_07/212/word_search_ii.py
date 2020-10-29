from typing import List


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        trie = {}
        end_mark = '#'
        tmp_mark = '@'
        for word in words:
            t = trie
            for w in word:
                if w not in t:
                    t[w] = {}
                t = t[w]
            t[end_mark] = {}

        res = []
        m, n = len(board), len(board[0])

        def dfs(x: int, y: int, node: dict, s: str) -> None:
            v = board[x][y]
            if v not in node:
                return

            s += v
            node = node[v]
            if end_mark in node:
                res.append(s)
                node.pop(end_mark)

            board[x][y] = tmp_mark
            for dx, dy in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                x_, y_ = x + dx, y + dy
                if 0 <= x_ < m and 0 <= y_ < n:
                    dfs(x_, y_, node, s)
            board[x][y] = v

        for i in range(m):
            for j in range(n):
                dfs(i, j, trie, "")
        return res

    def test(self):
        board = [
            ['o', 'a', 'a', 'n'],
            ['e', 't', 'a', 'e'],
            ['i', 'h', 'k', 'r'],
            ['i', 'f', 'l', 'v']
        ]

        words = ["oath", "pea", "eat", "rain", "eate", "eie", "a"]

        ans = self.findWords(board, words)
        print(ans)


if __name__ == "__main__":
    Solution().test()
