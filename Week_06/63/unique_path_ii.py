from typing import List

class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        if not obstacleGrid or not obstacleGrid[0] or obstacleGrid[0][0] == 1:
            return 0
        m, n = len(obstacleGrid), len(obstacleGrid[0])
        dp = [[0] * n for _ in range(m)]

        for i in range(m):
            for j in range(n):
                if obstacleGrid[i][j] == 0:
                    if i == j == 0:
                        v = 1
                    else:
                        v = 0
                        if i > 0 and obstacleGrid[i - 1][j] == 0:
                            v += dp[i - 1][j]
                        if j > 0 and obstacleGrid[i][j - 1] == 0:
                            v += dp[i][j - 1]
                    dp[i][j] = v
        return dp[m - 1][n - 1]
