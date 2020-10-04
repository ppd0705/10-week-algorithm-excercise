from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        cache = {0: 0}

        def dfs(n: int):
            if n in cache:
                return cache[n]
            if n < 0:
                return -1
            v = amount + 1
            for c in coins:
                if n >= c and dfs(n - c) >= 0:
                    v = min(dfs(n - c) + 1, v)
            cache[n] = v
            return v

        if dfs(amount) == amount + 1:
            return -1
        return dfs(amount)
