from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)
        if n < 2:
            return 0
        dp11 = -prices[0]
        dp10 = 0
        dp21 = -prices[1]
        dp20 = 0

        for p in prices[1:]:
            dp11 = max(dp11, -p)
            dp10 = max(dp10, dp11 + p)
            dp21 = max(dp21, dp10 - p)
            dp20 = max(dp20, dp21 + p)
        return dp20
