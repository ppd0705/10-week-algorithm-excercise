from typing import List


class Solution:
    def findCircleNum(self, M: List[List[int]]) -> int:
        n = len(M)
        res = 0
        parent = [-1] * n
        for i in range(n):
            for j in range(n):
                if i != j and M[i][j] == 1:
                    self.union(parent, i, j)
        for p in parent:
            if p == -1:
                res += 1

        return res

    def find(self, parent: List[int], i: int):
        if parent[i] == -1:
            return i
        return self.find(parent, parent[i])

    def union(self, parent: List[int], x: int, y: int):
        px = self.find(parent, x)
        py = self.find(parent, y)
        if px != py:
            parent[px] = py
