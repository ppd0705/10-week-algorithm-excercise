from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        def helper(nums: List[int]) -> int:
            pre = cur = 0
            for n in nums:
                pre, cur = cur, max(cur, pre + n)
            return cur

        if len(nums) == 0:
            return 0
        elif len(nums) == 1:
            return nums[0]
        return max(helper(nums[:-1]), helper(nums[1:]))
