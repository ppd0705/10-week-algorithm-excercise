from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        ans = []
        k = len(nums)
        option = [0] * k
        visited = [0] * k

        def dfs(idx: int):
            for i in range(k):
                if visited[i] == 0:
                    option[idx] = nums[i]
                    visited[i] = 1
                    if idx == k - 1:
                        ans.append(option.copy())
                    else:
                        dfs(idx + 1)
                    visited[i] = 0

        dfs(0)
        return ans
