from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        ans = []
        combination = [0] * k

        def dfs(start: int, idx: int):
            for i in range(start, n + 1 - k + idx + 1):
                combination[idx] = i
                if idx == k - 1:
                    ans.append(combination.copy())
                else:
                    dfs(i + 1, idx + 1)
        dfs(1, 0)
        return ans
