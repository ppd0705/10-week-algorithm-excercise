from typing import List


class NumMatrix:
    def __init__(self, matrix: List[List[int]]):
        m = len(matrix)
        n = 0
        if m > 0:
            n = len(matrix[0])

        self.dp = [[0] * (n + 1) for _ in range(m + 1)]

        for i in range(m):
            for j in range(n):
                self.dp[i + 1][j + 1] = matrix[i][j] + self.dp[i][j + 1] + self.dp[i + 1][j] - self.dp[i][j]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        return self.dp[row2 + 1][col2 + 1] - self.dp[row2 + 1][col1] - self.dp[row1][col2 + 1] + self.dp[row1][col1]
