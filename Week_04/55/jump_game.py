from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        right_most = 0
        n = len(nums)
        for i in range(0, n):
            if i <= right_most:
                right_most = max(i + nums[i], right_most)
            if right_most >= n - 1:
                return True
        return False
