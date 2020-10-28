from typing import List


class Solution:
    def canCross(self, stones: List[int]) -> bool:
        if stones[1] != 1:
            return False

        stones_set = set(stones)
        memo = {}

        def dfs(i: int, k: int) -> bool:
            if i == stones[-1]:
                return True
            if (i, k) in memo:
                return memo[(i, k)]
            res = False
            for j in range(k - 1, k + 2):
                if j > 0 and i + j in stones_set and dfs(i + j, j):
                    res = True
                    break
            memo[(i, k)] = res
            return res

        return dfs(0, 0)
