from typing import List


class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        m, n = len(grid) - 1, len(grid[0]) - 1
        if grid[0][0] == 1 or grid[m][n] == 1:
            return -1

        offsets = [
            (0, 1),
            (0, -1),
            (1, -1),
            (1, 0),
            (1, 1),
            (-1, -1),
            (-1, 0),
            (-1, 1),
        ]

        queue = [(0, 0)]
        res = 0
        while queue:
            res += 1
            l = len(queue)
            for i in range(l):
                if queue[i][0] == m and queue[i][1] == n:
                    return res
                for dx, dy in offsets:
                    x, y = queue[i][0] + dx, queue[i][1] + dy
                    if 0 <= x <= m and 0 <= y <= n and grid[x][y] == 0:
                        queue.append((x, y))
                        grid[x][y] = 1
            queue = queue[l:]
        return -1
