from typing import List


class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        if k == 0 or len(prices) < 2:
            return 0
        if 2 * k >= len(prices):
            return self.maxProfitWithoutLimit(prices)

        dp = [[0, 0] for _ in range(k + 1)]

        for i in range(len(prices)):
            for j in range(1, k + 1):
                if i == 0:
                    dp[j][1] = -prices[i]
                else:
                    dp[j][0] = max(dp[j][0], dp[j][1] + prices[i])
                    dp[j][1] = max(dp[j][1], dp[j - 1][0] - prices[i])
        return dp[k][0]

    def maxProfitWithoutLimit(self, prices: List[int]) -> int:
        res = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i - 1]:
                res += prices[i] - prices[i - 1]
        return res
