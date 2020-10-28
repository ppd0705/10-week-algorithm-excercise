from typing import List


class Solution:
    def canCross(self, stones: List[int]) -> bool:
        if stones[1] != 1:
            return False
        n = len(stones)

        dp = [[False] * n for _ in range(n)]
        dp[0][0] = True
        for i in range(1, n):
            for j in range(0, i):
                k = stones[i] - stones[j]
                if k < 1 or k > j + 1:
                    continue
                dp[i][k] = dp[j][k - 1] or dp[j][k] or dp[j][k + 1]
                if i == n - 1 and dp[i][k]:
                    return True
        return False
