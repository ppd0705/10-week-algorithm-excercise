from typing import List


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        ans = []
        len(nums)
        option = [0] * len(nums)
        elements = {}
        for n in nums:
            if n in elements:
                elements[n] += 1
            else:
                elements[n] = 1

        def dfs(idx: int):
            for k, v in elements.items():
                if v > 0:
                    option[idx] = k
                    elements[k] -= 1
                    if idx == len(nums) - 1:
                        ans.append(option.copy())
                    else:
                        dfs(idx + 1)
                    elements[k] += 1

        dfs(0)
        return ans
