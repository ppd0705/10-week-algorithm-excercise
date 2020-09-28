from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []

        def dfs(arr: List[int], idx: int):
            res.append(arr.copy())
            for i in range(idx, len(nums)):
                dfs(arr + [i])

        dfs([], 0)
        return res
