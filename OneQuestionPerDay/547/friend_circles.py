from typing import List


class Solution:
    def findCircleNum(self, M: List[List[int]]) -> int:
        n = len(M)
        res = 0
        visited = [False] * n

        def dfs(i: int):
            for j in range(n):
                if not visited[j] and M[i][j]:
                    visited[j] = True
                    dfs(j)

        for i in range(n):
            if not visited[i]:
                res += 1
                dfs(i)
        return res
