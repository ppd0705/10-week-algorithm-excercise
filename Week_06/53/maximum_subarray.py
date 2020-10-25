from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        curr = nums[0]
        res = curr

        for n in nums[1:]:
            curr = n + max(0, curr)
            res = max(res, curr)
        return res
