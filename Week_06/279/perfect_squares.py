import math


class Solution:
    def numSquares(self, n: int) -> int:
        dp = [n + 1] * (n + 1)
        dp[0] = 0

        squares = [i * i for i in range(1, int(math.sqrt(n)) + 1)]

        for i in range(1, n + 1):
            for j in squares:
                if i >= j:
                    dp[i] = min(dp[i], dp[i - j] + 1)
        return dp[-1]
