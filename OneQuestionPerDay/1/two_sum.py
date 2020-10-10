from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        cache = {}
        for i, n in enumerate(nums):
            m = target - n
            if m in cache:
                return [cache[m], i]
            cache[n] = i
        return []