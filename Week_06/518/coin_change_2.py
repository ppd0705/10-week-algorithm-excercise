from typing import List

class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        if not coins or not amount:
            return 0
        dp = [0] * (amount+1)
        dp[0] = 1

        for i in range(1, amount+1):
            for c in coins:
                if i >= c:
                    dp[i] += dp[i-c]
        return dp[amount]