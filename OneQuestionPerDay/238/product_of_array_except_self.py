from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        res = [0] * len(nums)
        k = 1
        for i in range(len(nums)):
            res[i] = k
            k *= nums[i]
        k = 1
        for i in range(len(nums) - 1, -1, -1):
            res[i] *= k
            k *= nums[i]
        return res
