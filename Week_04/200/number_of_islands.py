from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        m = len(grid)
        if m == 0:
            return 0
        n = len(grid[0])
        if n == 0:
            return 0
        res = 0

        def dfs(x, y):
            grid[x][y] = "0"

            if x - 1 >= 0 and grid[x][y] == "1":
                dfs(x - 1, y)
            if x + 1 < m and grid[x + 1][y] == "1":
                dfs(x + 1, y)

            if y - 1 >= 0 and grid[x][y - 1] == "1":
                dfs(x, y - 1)
            if y + 1 < n and grid[x][y + 1] == "1":
                dfs(x, y + 1)

        for i in range(m):
            for j in range(n):
                if grid[i][j] == "1":
                    res += 1
                    dfs(i, j)
        return res
