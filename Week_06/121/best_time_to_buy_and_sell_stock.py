from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0
        min_price = prices[0]
        res = 0
        for p in prices[1:]:
            if p > min_price:
                res = max(res, p - min_price)
            else:
                min_price = p
        return res
